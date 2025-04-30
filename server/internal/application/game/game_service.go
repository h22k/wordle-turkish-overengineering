package application

import (
	"context"

	"github.com/h22k/wordle-turkish-overengineering/server/internal/application/game/command"
	"github.com/h22k/wordle-turkish-overengineering/server/internal/application/game/query"
	domain "github.com/h22k/wordle-turkish-overengineering/server/internal/domain/game"
)

type GameService struct {
	makeGuessCommand *command.MakeGuessCommand
	newGameCommand   *command.NewGameCommand
	addWordCommand   *command.WordCommand

	gameQuery          *query.GameQuery
	randomVocableQuery *query.VocableQuery
	guessQuery         *query.GuessQuery
}

func NewGameService(
	makeGuessCommand *command.MakeGuessCommand,
	newGameCommand *command.NewGameCommand,
	addWordCommand *command.WordCommand,
	gameQuery *query.GameQuery,
	randomVocableQuery *query.VocableQuery,
	guessQuery *query.GuessQuery,
) *GameService {
	return &GameService{
		makeGuessCommand:   makeGuessCommand,
		newGameCommand:     newGameCommand,
		addWordCommand:     addWordCommand,
		gameQuery:          gameQuery,
		randomVocableQuery: randomVocableQuery,
		guessQuery:         guessQuery,
	}
}

func (gs GameService) MakeGuess(ctx context.Context, input MakeGuessInput) (command.MakeGuessResult, error) {
	game, err := gs.gameQuery.GetLastGame(ctx)

	if err != nil {
		return command.MakeGuessResult{}, err
	}

	guesses, err := gs.guessQuery.GetGameGuesses(ctx, game, input.SessionId)

	if err != nil {
		return command.MakeGuessResult{}, err
	}

	if err = game.SetGuesses(guesses); err != nil {
		return command.MakeGuessResult{}, err
	}

	return gs.makeGuessCommand.Execute(ctx, makeGuessInputToCommandInput(input, game))
}

func (gs GameService) GetGameGuesses(ctx context.Context, game domain.Game, sessionId string) ([]domain.WordGuess, error) {
	return gs.guessQuery.GetGameGuesses(ctx, game, sessionId)
}

func (gs GameService) CreateGame(ctx context.Context) (command.CreateGameResult, error) {
	word, err := gs.randomVocableQuery.GetDailyWord(ctx)

	if err != nil {
		return command.CreateGameResult{}, err
	}

	gameResult, err := gs.newGameCommand.Execute(ctx, word)

	if err != nil {
		return command.CreateGameResult{}, err
	}

	return gameResult, nil
}

func (gs GameService) AddWord(ctx context.Context, word domain.Word) error {
	return gs.addWordCommand.AddWord(ctx, word)
}

func (gs GameService) LastGame(ctx context.Context) (domain.Game, error) {
	return gs.gameQuery.GetLastGame(ctx)
}

func makeGuessInputToCommandInput(input MakeGuessInput, game domain.Game) command.MakeGuessInput {
	return command.MakeGuessInput{
		Guess:     domain.Word(input.Word),
		SessionId: input.SessionId,
		Game:      game,
	}
}
