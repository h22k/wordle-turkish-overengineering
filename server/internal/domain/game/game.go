package domain

import (
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
)

const (
	GameCacheTtl = time.Hour * 24
)

var (
	LengthIsIncorrectErr       = errors.New("length is incorrect")
	MaxWordGuessesExceededErr  = errors.New("max word guesses exceeded")
	AlreadyGuessedCorrectlyErr = errors.New("already guessed correctly")
)

var EmptyGame Game

type Word string

func (w Word) Len() uint8 {
	wstr := string(w)
	wstr = strings.ToLower(wstr)
	wstr = strings.Trim(wstr, " ")
	return uint8(len(wstr))
}

func (w Word) String() string {
	return string(w)
}

type Game struct {
	ID             uuid.UUID
	Word           Word
	WordGuesses    []WordGuess
	MaxWordGuesses uint8
}

func NewGame(word Word) Game {
	return NewGameWithId(word, uuid.New())
}

func NewGameWithId(word Word, id uuid.UUID) Game {
	return Game{
		ID:             id,
		Word:           word,
		WordGuesses:    make([]WordGuess, 0, word.Len()+1),
		MaxWordGuesses: word.Len() + 1,
	}
}

func (g *Game) MakeGuess(guess Word) (WordGuess, error) {
	if g.Word.Len() != guess.Len() {
		return WordGuess{}, LengthIsIncorrectErr
	}

	if g.guessedCorrectly() {
		return WordGuess{}, AlreadyGuessedCorrectlyErr
	}

	if g.guessExceeded() {
		return WordGuess{}, MaxWordGuessesExceededErr
	}

	wordGuess := NewWordGuess(g.Word, guess)
	g.WordGuesses = append(g.WordGuesses, wordGuess)

	return wordGuess, nil
}

func (g *Game) guessedCorrectly() bool {
	return len(g.WordGuesses) >= 1 && g.WordGuesses[len(g.WordGuesses)-1].IsCorrect()
}

func (g *Game) guessExceeded() bool {
	return len(g.WordGuesses) >= int(g.MaxWordGuesses)
}
