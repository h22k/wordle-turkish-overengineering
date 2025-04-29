package game

type ActiveGameResponse struct {
	MaxGuesses uint8 `json:"max_guesses"`
}

type GuessedWordResponse struct {
	Word    string       `json:"word"`
	Letters []LetterView `json:"letters"`
}
