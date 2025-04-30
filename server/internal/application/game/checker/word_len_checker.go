package checker

import (
	"context"

	domain "github.com/h22k/wordle-turkish-overengineering/server/internal/domain/game"
)

type WordLenChecker struct {
}

func NewWordLenChecker() *WordLenChecker {
	return &WordLenChecker{}
}

func (w WordLenChecker) Check(ctx context.Context, word domain.Word) (bool, error) {
	if word.Len() < 5 || word.Len() > 7 {
		return false, domain.LengthIsIncorrectErr
	}

	// we return false because we want to continue to the next checker
	// against the possibility of no next checker, we return nil as error
	return false, nil
}
