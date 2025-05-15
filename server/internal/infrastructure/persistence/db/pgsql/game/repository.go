package game

import (
	"context"

	"github.com/google/uuid"
	domain "github.com/h22k/wordle-turkish-overengineering/server/internal/domain/game"
	"github.com/h22k/wordle-turkish-overengineering/server/internal/infrastructure/persistence/db/pgsql/game/query"
)

type Repository struct {
	queries *query.Queries
}

func NewRepository(queries *query.Queries) *Repository {
	return &Repository{
		queries: queries,
	}
}

func (r Repository) Save(ctx context.Context, game domain.Game, wordId int32) error {
	_, err := r.queries.CreateGame(ctx, query.CreateGameParams{
		WordID:      wordId,
		MaxAttempts: int32(game.MaxWordGuesses),
		WordLength:  int32(game.Word.Len()),
	})

	return err
}

func (r Repository) GetLastGame(ctx context.Context) (domain.Game, error) {
	game, err := r.queries.GetActiveGame(ctx)

	if err != nil {
		return domain.EmptyGame, err
	}

	return domain.NewGameWithId(domain.Word(game.SecretWord), game.ID), nil
}

func (r Repository) MakeGameInactive(ctx context.Context, gameId uuid.UUID) error {
	_, err := r.queries.MakeGameInactive(ctx, gameId)

	return err
}
