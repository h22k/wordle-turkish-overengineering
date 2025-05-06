package bootstrap

import (
	"github.com/h22k/wordle-turkish-overengineering/server/internal/application/game/command"
	"github.com/h22k/wordle-turkish-overengineering/server/internal/application/game/query"
	domain "github.com/h22k/wordle-turkish-overengineering/server/internal/domain/game"
)

type db interface {
	gameRepository() domain.GameRepository
	guessRepository() domain.GuessRepository
	vocableRepository() domain.VocableRepository
}

type cache interface {
	gameCacheRepository() domain.GameCacheRepository
}

type usecase struct {
	addWordCommand   *command.WordCommand
	newGameCommand   *command.NewGameCommand
	makeGuessCommand *command.MakeGuessCommand

	gameQuery    *query.GameQuery
	vocableQuery *query.VocableQuery
	guessQuery   *query.GuessQuery
}

func initUseCases(db db, cache cache) *usecase {
	return &usecase{
		addWordCommand:   command.NewWordCommand(db.vocableRepository()),
		newGameCommand:   command.NewNewGameCommand(db.gameRepository(), cache.gameCacheRepository()),
		makeGuessCommand: command.NewMakeGuessCommand(db.guessRepository()),

		gameQuery:    query.NewGameQuery(db.gameRepository(), cache.gameCacheRepository()),
		vocableQuery: query.NewVocableQuery(db.vocableRepository()),
		guessQuery:   query.NewGuessQuery(db.guessRepository()),
	}
}

func (u usecase) GuessQuery() *query.GuessQuery {
	return u.guessQuery
}

func (u usecase) AddWordCommand() *command.WordCommand {
	return u.addWordCommand
}

func (u usecase) NewGameCommand() *command.NewGameCommand {
	return u.newGameCommand
}

func (u usecase) MakeGuessCommand() *command.MakeGuessCommand {
	return u.makeGuessCommand
}

func (u usecase) GameQuery() *query.GameQuery {
	return u.gameQuery
}

func (u usecase) VocableQuery() *query.VocableQuery {
	return u.vocableQuery
}
