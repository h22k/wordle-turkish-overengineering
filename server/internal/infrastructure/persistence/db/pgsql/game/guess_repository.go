package game

import (
	"context"

	"github.com/google/uuid"
	domain "github.com/h22k/wordle-turkish-overengineering/server/internal/domain/game"
	"github.com/h22k/wordle-turkish-overengineering/server/internal/infrastructure/persistence/db/pgsql/game/query"
)

type GuessRepository struct {
	queries *query.Queries
}

func (g GuessRepository) Save(ctx context.Context, guess domain.WordGuess, game domain.Game, sessionId string) error {
	guessCount, err := g.queries.GetGameGuessesCount(ctx, query.GetGameGuessesCountParams{
		GameID:    game.ID,
		SessionID: sessionId,
	})

	if err != nil {
		return err
	}

	// we assume that int64 type of guessCount is always between 0 and 255.
	// so we can safely cast it to uint8
	if uint8(guessCount) >= game.MaxWordGuesses {
		return domain.MaxWordGuessesExceededErr
	}

	_, err = g.queries.CreateGuess(ctx, query.CreateGuessParams{
		GameID:        game.ID,
		SessionID:     sessionId,
		Word:          guess.Guess.String(),
		AttemptNumber: int32(guessCount + 1), // same above
	})

	if err != nil {
		return err
	}

	return nil
}

func (g GuessRepository) FindByGameAndSessionId(ctx context.Context, gameId uuid.UUID, sessionID string) ([]domain.WordGuess, error) {
	game, err := g.queries.FindGameById(ctx, gameId)

	if err != nil {
		return nil, err
	}

	guesses, err := g.queries.GetGameGuesses(ctx, query.GetGameGuessesParams{
		GameID:    gameId,
		SessionID: sessionID,
	})

	if err != nil {
		return nil, err
	}

	var wordGuesses []domain.WordGuess
	for _, guess := range guesses {
		wordGuesses = append(wordGuesses, domain.NewWordGuess(domain.Word(game.SecretWord), domain.Word(guess.Word)))
	}

	return wordGuesses, nil
}
