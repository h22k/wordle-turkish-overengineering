package bootstrap

import domain "github.com/h22k/wordle-turkish-overengineering/server/internal/domain/game"

// validator TODO:: will be implemented later
type validator struct {
}

func newValidator() *validator {
	return &validator{}
}

func (v validator) validator() domain.WordValidator {
	return nil
}
