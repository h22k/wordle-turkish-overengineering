package redis

import (
	"context"
	"time"

	domain "github.com/h22k/wordle-turkish-overengineering/server/internal/domain/game"
)

// GameCacheRepository TODO:: empty implementation for now
type GameCacheRepository struct{}

func (g GameCacheRepository) Save(ctx context.Context, game domain.Game, ttl time.Duration) error {
	return nil
}

func (g GameCacheRepository) Get(ctx context.Context) (domain.Game, error) {
	return domain.EmptyGame, nil
}
