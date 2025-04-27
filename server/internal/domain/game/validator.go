package domain

import "context"

type WordValidator interface {
	Validate(ctx context.Context, word Word) error
}
