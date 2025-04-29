package bootstrap

import (
	application "github.com/h22k/wordle-turkish-overengineering/server/internal/application/game"
	"github.com/h22k/wordle-turkish-overengineering/server/internal/application/game/command"
	"github.com/h22k/wordle-turkish-overengineering/server/internal/application/game/query"
)

type ucase interface {
	AddWordCommand() *command.AddWordCommand
	NewGameCommand() *command.NewGameCommand
	MakeGuessCommand() *command.MakeGuessCommand

	GameQuery() *query.GameQuery
	RandomVocableQuery() *query.RandomVocableQuery
	GuessQuery() *query.GuessQuery
}

type appService struct {
	gameService application.GameService
}

func initService(uc ucase) *appService {
	return &appService{
		gameService: application.NewGameService(
			uc.MakeGuessCommand(),
			uc.NewGameCommand(),
			uc.AddWordCommand(),
			uc.GameQuery(),
			uc.RandomVocableQuery(),
			uc.GuessQuery(),
		),
	}
}
