package bootstrap

import (
	"context"
	"fmt"

	"github.com/grafana/pyroscope-go"
	"github.com/h22k/wordle-turkish-overengineering/server/config"
	"github.com/h22k/wordle-turkish-overengineering/server/internal/application/event"
	validator3 "github.com/h22k/wordle-turkish-overengineering/server/internal/application/validator"
	metrics "github.com/h22k/wordle-turkish-overengineering/server/internal/infrastructure/metric"
	"github.com/h22k/wordle-turkish-overengineering/server/internal/infrastructure/persistence/db/pgsql"
	validator2 "github.com/h22k/wordle-turkish-overengineering/server/internal/infrastructure/validator"
	"github.com/h22k/wordle-turkish-overengineering/server/internal/presentation/http/echo/game"
	"github.com/h22k/wordle-turkish-overengineering/server/internal/presentation/http/echo/middleware"
	game2 "github.com/h22k/wordle-turkish-overengineering/server/internal/presentation/http/game"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Application struct {
	ctx        context.Context
	echoApp    *echo.Echo
	db         *pgxpool.Pool
	dbListener *pgsql.EventListener
	appService *appService
	cfg        config.Config
	profiler   *pyroscope.Profiler
}

func InitApplication(ctx context.Context, cfg config.Config) *Application {
	pgPoolConn := must(initPostgresqlPoolConn(ctx, cfg))
	pgQuery := initPostgresQuery(pgPoolConn)
	pgDb := newPostgresDb(pgQuery)
	pgListener := pgsql.NewEventListener(must(initPostgresqlConn(ctx, cfg)))

	useCase := initUseCases(pgDb, newRedisCache())
	as := initService(useCase, initChecker(useCase))

	e := echo.New()
	e.Validator = validator2.NewValidator()

	e.Use(echoMiddleware.CORSWithConfig(echoMiddleware.CORSConfig{
		AllowOrigins: cfg.AllowOrigins,
		AllowMethods: []string{echo.GET, echo.POST, echo.OPTIONS},
	}))

	return &Application{
		appService: as,
		cfg:        cfg,
		echoApp:    e,
		db:         pgPoolConn,
		dbListener: pgListener,
		ctx:        ctx,
	}
}

func (a *Application) AppService() *appService {
	return a.appService
}

func (a *Application) SetRoutes() {
	go a.setProfiler()

	a.setMiddlewares()
	a.setMetric()

	v1 := a.echoApp.Group("/api/v1")
	v1.Use(echoMiddleware.Logger())

	gameRoute := v1.Group("/game")

	v := validator2.NewValidator()

	a.setGameRoutes(gameRoute, v)
}

func (a *Application) Run() error {
	return a.echoApp.Start(a.cfg.ServerPort)
}

func (a *Application) Close() {
	if err := a.echoApp.Shutdown(a.ctx); err != nil {
		a.echoApp.Logger.Fatal(err)
	}
	a.db.Close()

	if a.profiler != nil {
		a.profiler.Stop()
	}
}

func (a *Application) setMetric() {
	metrics.Init()

	a.echoApp.GET("/metrics", echo.WrapHandler(promhttp.Handler()))
}

func (a *Application) setMiddlewares() {
	a.echoApp.Use(middleware.MetricsMiddleware())
	a.echoApp.Use(middleware.ServerTimingMiddleware())
	a.echoApp.Use(middleware.IdentifierCookieMiddleware(a.cfg.CookieDomain, a.cfg.IsProd()))
	a.echoApp.Use(echoMiddleware.RequestID())
	a.echoApp.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Response().Header().Set(echo.HeaderAccessControlAllowCredentials, "true")
			return next(c)
		}
	})
}

func (a *Application) setGameRoutes(gameRoute *echo.Group, v validator3.InputValidator) {
	gameHandler := game.NewHandler(game2.NewService(a.appService.gameService, a.appService.wordChecker), v)

	gameBroker := event.NewBroker()
	gameDispatcher := event.NewDispatcher(a.dbListener, gameBroker)

	go gameDispatcher.Start(a.ctx, "game_created")

	gameRoute.GET("/game", gameHandler.GetGame())
	gameRoute.POST("/guess", middleware.ValidateRequest(gameHandler.MakeGuess))
	gameRoute.GET("/events", gameHandler.Sse(gameBroker))
}

func (a *Application) setProfiler() {
	profiler, err := pyroscope.Start(pyroscope.Config{
		ApplicationName: a.cfg.AppName,
		ServerAddress:   a.cfg.PyroscopeUrl,
		Logger:          pyroscope.StandardLogger,
		ProfileTypes: []pyroscope.ProfileType{
			pyroscope.ProfileCPU,
			pyroscope.ProfileAllocObjects,
			pyroscope.ProfileAllocSpace,
			pyroscope.ProfileInuseObjects,
			pyroscope.ProfileInuseSpace,
			pyroscope.ProfileGoroutines,
		},
	})

	if err != nil {
		panic(fmt.Sprintf("Failed to start pyroscope profiler: %v", err))
	}

	a.profiler = profiler
}
