package game

import (
	"context"

	application "github.com/h22k/wordle-turkish-overengineering/server/internal/application/game"
	domain "github.com/h22k/wordle-turkish-overengineering/server/internal/domain/game"
)

type Service struct {
	gameService *application.GameService
	wordChecker *domain.WordCheckerChain
}

func NewService(gameService *application.GameService, checker *domain.WordCheckerChain) *Service {
	return &Service{
		gameService: gameService,
		wordChecker: checker,
	}
}

func (s *Service) GetGameInfo(ctx context.Context, sessionId string) (domain.Game, error) {
	game, err := s.gameService.LastGame(ctx)
	if err != nil {
		return domain.EmptyGame, err
	}

	guesses, err := s.gameService.GetGameGuesses(ctx, game, sessionId)
	if err != nil {
		return domain.EmptyGame, err
	}

	if err = game.SetGuesses(guesses); err != nil {
		return domain.EmptyGame, err
	}

	return game, nil
}

func (s *Service) MakeGuess(ctx context.Context, sessionId string, guess string) (domain.WordGuess, error) {
	err := s.wordChecker.Check(ctx, domain.Word(guess))
	if err != nil {
		return domain.WordGuess{}, err
	}

	result, err := s.gameService.MakeGuess(ctx, application.MakeGuessInput{
		SessionId: sessionId,
		Word:      guess,
	})
	if err != nil {
		return domain.WordGuess{}, err
	}

	return result.Guess, nil
}
