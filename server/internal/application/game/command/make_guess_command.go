package command

import (
	"context"

	domain "github.com/h22k/wordle-turkish-overengineering/server/internal/domain/game"
)

type MakeGuessCommand struct {
	guessRepository domain.GuessRepository
}

func NewMakeGuessCommand(guessRepo domain.GuessRepository) *MakeGuessCommand {
	return &MakeGuessCommand{
		guessRepository: guessRepo,
	}
}

func (mgc MakeGuessCommand) Execute(ctx context.Context, input MakeGuessInput) (MakeGuessResult, error) {
	guessWord := input.Guess
	guess, err := input.Game.MakeGuess(guessWord)

	if err != nil {
		return MakeGuessResult{}, err
	}

	err = mgc.guessRepository.Save(ctx, guess, input.Game, input.SessionId)

	if err != nil {
		return MakeGuessResult{}, err
	}

	return MakeGuessResult{Guess: guess}, nil
}
