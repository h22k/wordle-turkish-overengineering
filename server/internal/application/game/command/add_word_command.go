package command

import (
	"context"

	domain "github.com/h22k/wordle-turkish-overengineering/server/internal/domain/game"
)

type WordCommand struct {
	WordRepo domain.VocableRepository
}

func NewWordCommand(wordRepo domain.VocableRepository) *WordCommand {
	return &WordCommand{
		WordRepo: wordRepo,
	}
}

func (awc WordCommand) AddWord(ctx context.Context, word domain.Word) error {
	if err := awc.WordRepo.Save(ctx, domain.NewVocable(word)); err != nil {
		return err
	}

	return nil
}
