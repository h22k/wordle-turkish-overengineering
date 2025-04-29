package game

import (
	"context"

	application "github.com/h22k/wordle-turkish-overengineering/server/internal/application/game"
	domain "github.com/h22k/wordle-turkish-overengineering/server/internal/domain/game"
)

type Service struct {
	gameService application.GameService
}

func NewService(gameService application.GameService) *Service {
	return &Service{
		gameService: gameService,
	}
}

func (s *Service) GetGameInfo(ctx context.Context, sessionId string) (domain.Game, error) {
	game, err := s.gameService.LastGame(ctx)

	if err != nil {
		return domain.EmptyGame, err
	}

	return game, nil
}

func (s *Service) MakeGuess(ctx context.Context, sessionId string, guess string) (domain.WordGuess, error) {
	result, err := s.gameService.MakeGuess(ctx, application.MakeGuessInput{
		SessionId: sessionId,
		Word:      guess,
	})

	if err != nil {
		return domain.WordGuess{}, err
	}

	return result.Guess, nil
}
