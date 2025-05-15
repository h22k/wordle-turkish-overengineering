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

func (rvq VocableQuery) GetDailyWord(ctx context.Context) (domain.Word, int32, error) {
	vocable, id, err := rvq.vocableRepository.FindRandom(ctx)

	if err != nil {
		return "", 0, err
	}

	return vocable, id, nil
}

func (rvq VocableQuery) IsWordAcceptable(ctx context.Context, word domain.Word) (bool, error) {
	return rvq.vocableRepository.IsWordExists(ctx, word)
}
