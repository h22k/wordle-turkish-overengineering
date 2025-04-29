package command

import (
	"context"

	domain "github.com/h22k/wordle-turkish-overengineering/server/internal/domain/game"
)

type AddWordCommand struct {
	WordRepo domain.VocableRepository
}

func NewAddWordCommand(wordRepo domain.VocableRepository) *AddWordCommand {
	return &AddWordCommand{
		WordRepo: wordRepo,
	}
}

func (awc AddWordCommand) Execute(ctx context.Context, word domain.Word) error {
	if err := awc.WordRepo.Save(ctx, domain.NewVocable(word)); err != nil {
		return err
	}

	return nil
}
