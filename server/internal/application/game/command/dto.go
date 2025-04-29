package command

import (
	"github.com/google/uuid"
	domain "github.com/h22k/wordle-turkish-overengineering/server/internal/domain/game"
)

type MakeGuessInput struct {
	Guess     domain.Word
	Game      domain.Game
	SessionId string
}

type MakeGuessResult struct {
	Guess domain.WordGuess
}

type CreateGameResult struct {
	ID               uuid.UUID
	WordLength       uint8
	MaxGuessAttempts uint8
}
