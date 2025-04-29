package domain

import (
	"errors"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/google/uuid"
)

const (
	GameCacheTtl = time.Hour * 24
)

var (
	LengthIsIncorrectErr       = errors.New("length is incorrect")
	MaxWordGuessesExceededErr  = errors.New("max word guesses exceeded")
	AlreadyGuessedCorrectlyErr = errors.New("already guessed correctly")
	SameGuessErr               = errors.New("same guess")
)

var EmptyGame Game

type Word string

func (w Word) Len() uint8 {
	wstr := string(w)
	wstr = strings.ToLower(wstr)
	wstr = strings.Trim(wstr, " ")
	return uint8(utf8.RuneCountInString(wstr))
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

	if g.GuessedCorrectly() {
		return WordGuess{}, AlreadyGuessedCorrectlyErr
	}

	if g.GuessExceeded() {
		return WordGuess{}, MaxWordGuessesExceededErr
	}

	for _, g := range g.WordGuesses {
		if g.Guess.String() == guess.String() {
			return WordGuess{}, SameGuessErr
		}
	}

	wordGuess := NewWordGuess(g.Word, guess)
	g.WordGuesses = append(g.WordGuesses, wordGuess)

	return wordGuess, nil
}

func (g *Game) SetGuesses(guesses []WordGuess) error {
	if len(guesses) > int(g.MaxWordGuesses) {
		return MaxWordGuessesExceededErr
	}

	for _, guess := range guesses {
		if g.Word.Len() != guess.Guess.Len() {
			return LengthIsIncorrectErr
		}
	}

	g.WordGuesses = guesses

	return nil
}

func (g *Game) GuessedCorrectly() bool {
	return len(g.WordGuesses) >= 1 && g.WordGuesses[len(g.WordGuesses)-1].IsCorrect()
}

func (g *Game) GuessExceeded() bool {
	return len(g.WordGuesses) >= int(g.MaxWordGuesses)
}
