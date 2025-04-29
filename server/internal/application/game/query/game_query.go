package query

import (
	"context"

	domain "github.com/h22k/wordle-turkish-overengineering/server/internal/domain/game"
)

type GameQuery struct {
	GameRepo      domain.GameRepository
	GameCacheRepo domain.GameCacheRepository
}

func NewGameQuery(gameRepo domain.GameRepository, gameCacheRepo domain.GameCacheRepository) *GameQuery {
	return &GameQuery{
		GameRepo:      gameRepo,
		GameCacheRepo: gameCacheRepo,
	}
}

func (gq *GameQuery) GetLastGame(ctx context.Context) (domain.Game, error) {
	if gq.GameCacheRepo == nil {
		return gq.lastGameFromDb(ctx)
	}

	game, err := gq.lastGameFromCache(ctx)

	if err == nil {
		return game, nil
	}

	return gq.lastGameFromDb(ctx)
}

func (gq *GameQuery) lastGameFromCache(ctx context.Context) (domain.Game, error) {
	return gq.GameCacheRepo.Get(ctx)
}

func (gq *GameQuery) lastGameFromDb(ctx context.Context) (domain.Game, error) {
	return gq.GameRepo.GetLastGame(ctx)
}
