package bootstrap

import (
	"context"
	"time"

	"github.com/h22k/wordle-turkish-overengineering/server/config"
	domain "github.com/h22k/wordle-turkish-overengineering/server/internal/domain/game"
	"github.com/h22k/wordle-turkish-overengineering/server/internal/infrastructure/persistence/db/pgsql/game"
	"github.com/h22k/wordle-turkish-overengineering/server/internal/infrastructure/persistence/db/pgsql/game/query"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type postgresDb struct {
	gameRepo    domain.GameRepository
	guessRepo   domain.GuessRepository
	vocableRepo domain.VocableRepository
}

func (p postgresDb) gameRepository() domain.GameRepository {
	return p.gameRepo
}

func (p postgresDb) guessRepository() domain.GuessRepository {
	return p.guessRepo
}

func (p postgresDb) vocableRepository() domain.VocableRepository {
	return p.vocableRepo
}

func newPostgresDb(q *query.Queries) *postgresDb {
	return &postgresDb{
		gameRepo:    game.NewRepository(q),
		guessRepo:   game.NewGuessRepository(q),
		vocableRepo: game.NewVocableRepository(q),
	}
}

func initPostgresqlConn(ctx context.Context, cfg config.Config) (*pgx.Conn, error) {
	pgConn, err := pgx.Connect(ctx, cfg.DbUrl)

	if err != nil {
		return nil, err
	}

	if err = pgConn.Ping(ctx); err != nil {
		return nil, err
	}

	return pgConn, nil
}

func initPostgresqlPoolConn(ctx context.Context, cfg config.Config) (*pgxpool.Pool, error) {
	pgConn, err := pgxpool.ParseConfig(cfg.DbUrl)

	if err != nil {
		return nil, err
	}

	pgConn.MaxConns = cfg.MaxDbConns
	pgConn.MinConns = cfg.MinDbConns
	pgConn.MaxConnLifetime = cfg.MaxDbConnLifeTime * time.Minute
	pgConn.MaxConnIdleTime = cfg.MaxDbIdleTime

	pool, err := pgxpool.NewWithConfig(ctx, pgConn)

	if err != nil {
		return nil, err
	}

	return pool, nil
}

func must[T any](val T, err error) T {
	if err != nil {
		panic(err)
	}

	return val
}

func initPostgresQuery(db query.DBTX) *query.Queries {
	return query.New(db)
}
