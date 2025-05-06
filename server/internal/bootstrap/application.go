package bootstrap

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/gofiber/fiber/v2/middleware/encryptcookie"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/grafana/pyroscope-go"
	"github.com/h22k/wordle-turkish-overengineering/server/config"
	validator3 "github.com/h22k/wordle-turkish-overengineering/server/internal/application/validator"
	validator2 "github.com/h22k/wordle-turkish-overengineering/server/internal/infrastructure/validator"
	metrics "github.com/h22k/wordle-turkish-overengineering/server/internal/metric"
	"github.com/h22k/wordle-turkish-overengineering/server/internal/presentation/http/fiber/game"
	"github.com/h22k/wordle-turkish-overengineering/server/internal/presentation/http/fiber/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Application struct {
	fiberApp *fiber.App

	db *pgxpool.Pool

	appService *appService

	cfg config.Config

	profiler *pyroscope.Profiler
}

func InitApplication(ctx context.Context, cfg config.Config) *Application {
	pgPoolConn := must(initPostgresql(ctx, cfg))
	pgQuery := initPostgresQuery(pgPoolConn)
	pgDb := newPostgresDb(pgQuery)

	useCase := initUseCases(pgDb, newRedisCache())
	as := initService(useCase, initChecker(useCase))

	return &Application{
		appService: as,
		cfg:        cfg,
		fiberApp:   fiber.New(fiber.Config{AppName: cfg.AppName}),
		db:         pgPoolConn,
	}
}

func (a *Application) AppService() *appService {
	return a.appService
}

func (a *Application) SetRoutes() {
	go a.setProfiler()

	a.setMiddlewares()
	a.setMetric()

	v1 := a.fiberApp.Group("/api/v1")
	v1.Use(logger.New())

	gameRoute := v1.Group("/game")

	v := validator2.NewValidator()

	a.setGameRoutes(gameRoute, v)
}

func (a *Application) Run() error {
	return a.fiberApp.Listen(a.cfg.ServerPort)
}

func (a *Application) Close() {
	a.fiberApp.Shutdown()
	a.db.Close()

	if a.profiler != nil {
		a.profiler.Stop()
	}
}

func (a *Application) setMetric() {
	metrics.Init()

	a.fiberApp.Get("/metrics", adaptor.HTTPHandler(promhttp.Handler()))
}

func (a *Application) setMiddlewares() {
	a.fiberApp.Hooks().OnRoute(func(route fiber.Route) error {
		fmt.Println(route.Path)
		return nil
	})

	a.fiberApp.Use(middleware.MetricsMiddleware())

	a.fiberApp.Use(middleware.ServerTimingMiddleware())

	a.fiberApp.Use(encryptcookie.New(encryptcookie.Config{
		Key: a.cfg.AppKey,
	}))

	a.fiberApp.Use(middleware.IdentifierCookieMiddleware())

	a.fiberApp.Use(requestid.New())
}

func (a *Application) setGameRoutes(gameRoute fiber.Router, v validator3.InputValidator) {
	gameHandler := game.NewHandler(game.NewService(a.appService.gameService, a.appService.wordChecker), v)

	gameRoute.Get("/game", gameHandler.GetGame())
	gameRoute.Post("/guess", gameHandler.MakeGuess())
}

func (a *Application) setProfiler() {
	profiler, err := pyroscope.Start(pyroscope.Config{
		ApplicationName: a.cfg.AppName,

		ServerAddress: a.cfg.PyroscopeUrl,

		Logger: pyroscope.StandardLogger,

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
