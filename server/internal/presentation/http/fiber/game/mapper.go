package game

import domain "github.com/h22k/wordle-turkish-overengineering/server/internal/domain/game"

func lettersToView(letters []domain.Letter) []LetterView {
	result := make([]LetterView, len(letters))
	for i, l := range letters {
		result[i] = LetterView{
			Char:   string(l.Char),
			Status: string(l.Status),
		}
	}
	return result
}

func guessesToResponse(guesses []domain.WordGuess) []GuessedWordResponse {
	result := make([]GuessedWordResponse, len(guesses))
	for i, g := range guesses {
		result[i] = GuessedWordResponse{
			Word:    g.Guess.String(),
			Letters: lettersToView(g.Letters),
		}
	}
	return result
}
