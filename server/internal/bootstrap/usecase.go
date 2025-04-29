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

type wordValidator interface {
	validator() domain.WordValidator
}

type usecase struct {
	addWordCommand   *command.AddWordCommand
	newGameCommand   *command.NewGameCommand
	makeGuessCommand *command.MakeGuessCommand

	gameQuery          *query.GameQuery
	randomVocableQuery *query.RandomVocableQuery
}

func initUseCases(db db, cache cache, wv wordValidator) *usecase {
	return &usecase{
		addWordCommand:   command.NewAddWordCommand(db.vocableRepository()),
		newGameCommand:   command.NewNewGameCommand(db.gameRepository(), cache.gameCacheRepository(), wv.validator()),
		makeGuessCommand: command.NewMakeGuessCommand(db.guessRepository(), db.gameRepository()),

		gameQuery:          query.NewGameQuery(db.gameRepository(), cache.gameCacheRepository()),
		randomVocableQuery: query.NewRandomVocableQuery(db.vocableRepository()),
	}
}

func (u usecase) AddWordCommand() *command.AddWordCommand {
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

func (u usecase) RandomVocableQuery() *query.RandomVocableQuery {
	return u.randomVocableQuery
}
