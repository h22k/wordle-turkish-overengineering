package db

import (
	"context"

	"github.com/google/uuid"
	domain "github.com/h22k/wordle-turkish-overengineering/server/internal/domain/game"
)

// GuessRepository This is a wrapper for the guess repository
type GuessRepository struct {
	Gr domain.GuessRepository
}

func (g GuessRepository) FindByGameAndSessionId(ctx context.Context, gameId uuid.UUID, sessionID string) ([]domain.WordGuess, error) {
	return g.Gr.FindByGameAndSessionId(ctx, gameId, sessionID)
}

func (g GuessRepository) Save(ctx context.Context, guess domain.WordGuess, game domain.Game, sessionId string) error {
	return g.Gr.Save(ctx, guess, game, sessionId)
}
