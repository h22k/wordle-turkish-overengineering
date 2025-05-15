package game

type MakeGuessRequest struct {
	Guess string `json:"guess" validate:"required,max=7,min=5"`
}

type LetterView struct {
	Char   string `json:"char"`
	Status string `json:"status"`
}

type SSEEvent struct {
	EventName string `json:"event_name"`
	Payload   string `json:"payload"`
}
