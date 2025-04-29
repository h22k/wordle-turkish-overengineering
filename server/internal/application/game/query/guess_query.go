package query

import (
	"context"

	domain "github.com/h22k/wordle-turkish-overengineering/server/internal/domain/game"
)

type GuessQuery struct {
	GuessRepo domain.GuessRepository
}

func NewGuessQuery(guessRepo domain.GuessRepository) *GuessQuery {
	return &GuessQuery{
		GuessRepo: guessRepo,
	}
}

func (gq *GuessQuery) GetGameGuesses(ctx context.Context, game domain.Game, sessionId string) ([]domain.WordGuess, error) {
	return gq.GuessRepo.FindByGameAndSessionId(ctx, game.ID, sessionId)
}
