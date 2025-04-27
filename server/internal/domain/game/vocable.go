package domain

import "github.com/google/uuid"

type Vocable struct {
	ID   uuid.UUID
	Word Word

	IsActive bool
}

func NewVocable(word Word) Vocable {
	return NewVocableWithID(uuid.Must(uuid.NewRandom()), word)
}

func NewVocableWithID(id uuid.UUID, word Word) Vocable {
	return Vocable{
		ID:       id,
		Word:     word,
		IsActive: true,
	}
}
