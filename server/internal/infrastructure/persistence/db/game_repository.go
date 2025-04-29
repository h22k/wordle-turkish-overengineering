package db

import (
	"context"

	domain "github.com/h22k/wordle-turkish-overengineering/server/internal/domain/game"
)

// GameRepository This is a wrapper for the game repository
type GameRepository struct {
	Repo domain.GameRepository
}

func (g GameRepository) Save(ctx context.Context, game domain.Game) error {
	return g.Repo.Save(ctx, game)
}

func (g GameRepository) GetLastGame(ctx context.Context) (domain.Game, error) {
	return g.Repo.GetLastGame(ctx)
}
