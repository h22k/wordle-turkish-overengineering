package query

import (
	"context"

	domain "github.com/h22k/wordle-turkish-overengineering/server/internal/domain/game"
)

type RandomVocableQuery struct {
	VocableRepository domain.VocableRepository
}

func (rvq RandomVocableQuery) GetDailyWord(ctx context.Context) (domain.Word, error) {
	vocable, err := rvq.VocableRepository.FindRandom(ctx)

	if err != nil {
		return "", err
	}

	return vocable, nil
}
