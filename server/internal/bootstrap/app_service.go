package bootstrap

import (
	application "github.com/h22k/wordle-turkish-overengineering/server/internal/application/game"
	"github.com/h22k/wordle-turkish-overengineering/server/internal/application/game/command"
	"github.com/h22k/wordle-turkish-overengineering/server/internal/application/game/query"
	domain "github.com/h22k/wordle-turkish-overengineering/server/internal/domain/game"
)

type ucase interface {
	AddWordCommand() *command.WordCommand
	NewGameCommand() *command.NewGameCommand
	MakeGuessCommand() *command.MakeGuessCommand

	MakeGameInactiveCommand() *command.MakeGameInactiveCommand

	GameQuery() *query.GameQuery
	VocableQuery() *query.VocableQuery
	GuessQuery() *query.GuessQuery
}

type appService struct {
	gameService *application.GameService

	wordChecker *domain.WordCheckerChain
}

func initService(uc ucase, wordCheckerChain *domain.WordCheckerChain) *appService {
	return &appService{
		gameService: application.NewGameService(
			uc.MakeGuessCommand(),
			uc.NewGameCommand(),
			uc.AddWordCommand(),
			uc.MakeGameInactiveCommand(),
			uc.GameQuery(),
			uc.VocableQuery(),
			uc.GuessQuery(),
		),
		wordChecker: wordCheckerChain,
	}
}

func (as appService) GameService() *application.GameService {
	return as.gameService
}
