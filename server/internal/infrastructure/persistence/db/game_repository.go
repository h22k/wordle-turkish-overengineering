package db

import (
	"context"

	"github.com/google/uuid"
	domain "github.com/h22k/wordle-turkish-overengineering/server/internal/domain/game"
)

// GameRepository This is a wrapper for the game repository
type GameRepository struct {
	Repo domain.GameRepository
}

func (g GameRepository) Save(ctx context.Context, game domain.Game, wordId int32) error {
	return g.Repo.Save(ctx, game, wordId)
}

func (g GameRepository) GetLastGame(ctx context.Context) (domain.Game, error) {
	return g.Repo.GetLastGame(ctx)
}

func (g GameRepository) MakeGameInactive(ctx context.Context, gameId uuid.UUID) error {
	return g.Repo.MakeGameInactive(ctx, gameId)
}
