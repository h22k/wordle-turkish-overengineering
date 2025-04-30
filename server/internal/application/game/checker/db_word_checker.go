package checker

import (
	"context"

	"github.com/h22k/wordle-turkish-overengineering/server/internal/application/game/query"
	domain "github.com/h22k/wordle-turkish-overengineering/server/internal/domain/game"
)

type DatabaseWordChecker struct {
	vq *query.VocableQuery
}

func NewDatabaseWordChecker(vq *query.VocableQuery) *DatabaseWordChecker {
	return &DatabaseWordChecker{
		vq: vq,
	}
}

func (w DatabaseWordChecker) Check(ctx context.Context, word domain.Word) (bool, error) {
	acceptable, err := w.vq.IsWordAcceptable(ctx, word)

	if err != nil {
		return false, err
	}

	return acceptable, nil
}
