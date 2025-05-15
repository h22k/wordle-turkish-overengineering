package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type GameRepository interface {
	Save(ctx context.Context, game Game, wordId int32) error
	GetLastGame(ctx context.Context) (Game, error)
	MakeGameInactive(ctx context.Context, gameId uuid.UUID) error
}

type GameCacheRepository interface {
	Save(ctx context.Context, game Game, ttl time.Duration) error
	Get(ctx context.Context) (Game, error)
}

type GuessRepository interface {
	FindByGameAndSessionId(ctx context.Context, gameId uuid.UUID, sessionID string) ([]WordGuess, error)
	Save(ctx context.Context, guess WordGuess, game Game, sessionId string) error
}

type VocableRepository interface {
	FindRandom(ctx context.Context) (Word, int32, error)
	Update(ctx context.Context, vocable Vocable) error
	Save(ctx context.Context, vocable Vocable) error
	FindByWord(ctx context.Context, word Word) (Vocable, error)
	IsWordExists(ctx context.Context, word Word) (bool, error)
}
