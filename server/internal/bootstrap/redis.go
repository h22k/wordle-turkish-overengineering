package bootstrap

import domain "github.com/h22k/wordle-turkish-overengineering/server/internal/domain/game"

// redisCache TODO:: will be implemented later
type redisCache struct {
}

func newRedisCache() *redisCache {
	return &redisCache{}
}

func (r redisCache) gameCacheRepository() domain.GameCacheRepository {
	return nil
}
