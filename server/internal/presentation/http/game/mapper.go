package game

import (
	"github.com/h22k/wordle-turkish-overengineering/server/internal/domain/event"
	domain "github.com/h22k/wordle-turkish-overengineering/server/internal/domain/game"
)

func LettersToView(letters []domain.Letter) []LetterView {
	result := make([]LetterView, len(letters))
	for i, l := range letters {
		result[i] = LetterView{
			Char:   string(l.Char),
			Status: string(l.Status),
		}
	}
	return result
}

func GuessesToResponse(guesses []domain.WordGuess) []GuessedWordResponse {
	result := make([]GuessedWordResponse, len(guesses))
	for i, g := range guesses {
		result[i] = GuessedWordResponse{
			Word:    g.Guess.String(),
			Letters: LettersToView(g.Letters),
		}
	}
	return result
}

func EventToSseEvent(event event.Event) SSEEvent {
	return SSEEvent{
		EventName: event.Name(),
		Payload:   event.Payload(),
	}
}
