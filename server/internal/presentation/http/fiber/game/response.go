package game

type ActiveGameResponse struct {
	MaxGuesses     uint8                 `json:"max_guesses"`
	IsGameFinished bool                  `json:"is_game_finished"`
	Guesses        []GuessedWordResponse `json:"guesses"`
}

type GuessedWordResponse struct {
	Word    string       `json:"word"`
	Letters []LetterView `json:"letters"`
}
