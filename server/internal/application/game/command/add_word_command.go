package command

import (
	"context"

	domain "github.com/h22k/wordle-turkish-overengineering/server/internal/domain/game"
)

type WordCommand struct {
	wordRepo domain.VocableRepository
}

func NewWordCommand(wordRepo domain.VocableRepository) *WordCommand {
	return &WordCommand{
		wordRepo: wordRepo,
	}
}

func (awc WordCommand) AddWord(ctx context.Context, word domain.Word) error {
	if err := awc.wordRepo.Save(ctx, domain.NewVocable(word)); err != nil {
		return err
	}

	return nil
}
