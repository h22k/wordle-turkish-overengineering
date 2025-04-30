package query

import (
	"context"

	domain "github.com/h22k/wordle-turkish-overengineering/server/internal/domain/game"
)

type VocableQuery struct {
	vocableRepository domain.VocableRepository
}

func NewVocableQuery(vocableRepo domain.VocableRepository) *VocableQuery {
	return &VocableQuery{
		vocableRepository: vocableRepo,
	}
}

func (rvq VocableQuery) GetDailyWord(ctx context.Context) (domain.Word, error) {
	vocable, err := rvq.vocableRepository.FindRandom(ctx)

	if err != nil {
		return "", err
	}

	return vocable, nil
}

func (rvq VocableQuery) IsWordAcceptable(ctx context.Context, word domain.Word) (bool, error) {
	vocable, err := rvq.vocableRepository.FindByWord(ctx, word)

	if err != nil || vocable.Word.Len() == 0 {
		return false, nil // TODO:: find better solution
	}

	return true, nil
}
