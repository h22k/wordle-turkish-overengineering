package command

import (
	"context"

	domain "github.com/h22k/wordle-turkish-overengineering/server/internal/domain/game"
)

type NewGameCommand struct {
	gameRepo      domain.GameRepository
	gameCacheRepo domain.GameCacheRepository
}

func NewNewGameCommand(gameRepo domain.GameRepository, gameCacheRepo domain.GameCacheRepository) *NewGameCommand {
	return &NewGameCommand{
		gameRepo:      gameRepo,
		gameCacheRepo: gameCacheRepo,
	}
}

func (ngc NewGameCommand) Execute(ctx context.Context, word domain.Word, id int32) (CreateGameResult, error) {
	game := domain.NewGame(word)
	err := ngc.gameRepo.Save(ctx, game, id)

	if err != nil {
		return CreateGameResult{}, err
	}

	if ngc.gameCacheRepo != nil {
		err = ngc.gameCacheRepo.Save(ctx, game, domain.GameCacheTtl)
	}

	// whether err is nil or not, we return it.
	// if err is not nil, it means that the game was not saved in the cache
	return CreateGameResult{
		ID:               game.ID,
		WordLength:       game.Word.Len(),
		MaxGuessAttempts: game.MaxWordGuesses,
	}, err
}
