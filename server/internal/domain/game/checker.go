package domain

import (
	"context"
	"errors"
)

var WordNotAcceptableErr = errors.New("word not acceptable")

type WordChecker interface {
	Check(ctx context.Context, word Word) (bool, error)
}

type WordCheckerChain struct {
	checkers []WordChecker
}

func NewWordCheckerChain(checkers ...WordChecker) *WordCheckerChain {
	return &WordCheckerChain{
		checkers: checkers,
	}
}

func (w WordCheckerChain) Check(ctx context.Context, word Word) error {
	for _, checker := range w.checkers {
		res, err := checker.Check(ctx, word)

		if err != nil {
			return err
		}

		if res {
			return nil
		}
	}

	return WordNotAcceptableErr
}
