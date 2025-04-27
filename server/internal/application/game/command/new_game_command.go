package command

import (
	"context"

	domain "github.com/h22k/wordle-turkish-overengineering/server/internal/domain/game"
)

type NewGameCommand struct {
	GameRepo      domain.GameRepository
	GameCacheRepo domain.GameCacheRepository
	WordValidator domain.WordValidator
}

func (ngc NewGameCommand) Execute(ctx context.Context, word domain.Word) (CreateGameResult, error) {
	if err := ngc.WordValidator.Validate(ctx, word); err != nil {
		return CreateGameResult{}, err
	}

	game := domain.NewGame(word)
	err := ngc.GameRepo.Save(ctx, game)

	if err != nil {
		return CreateGameResult{}, err
	}

	err = ngc.GameCacheRepo.Save(ctx, game, domain.GameCacheTtl)

	// whether err is nil or not, we return it. if err is not nil, it means that the game was not saved in the cache
	return CreateGameResult{
		ID:               game.ID,
		WordLength:       game.Word.Len(),
		MaxGuessAttempts: game.MaxWordGuesses,
	}, err
}
