package query

import (
	"context"

	domain "github.com/h22k/wordle-turkish-overengineering/server/internal/domain/game"
)

type GameQuery struct {
	GameRepo      domain.GameRepository
	GameCacheRepo domain.GameCacheRepository
}

func (gq *GameQuery) GetLastGame(ctx context.Context) (domain.Game, error) {
	game, err := gq.GameCacheRepo.Get(ctx)

	if err == nil {
		return game, err
	}

	game, err = gq.GameRepo.GetLastGame(ctx)

	if err != nil {
		return domain.EmptyGame, err
	}

	return game, nil
}
