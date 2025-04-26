package domain

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGame_MakeGuess(t *testing.T) {
	type fields struct {
		Id             uuid.UUID
		Word           Word
		WordGuesses    []WordGuess
		MaxWordGuesses uint8
	}
	type args struct {
		guess Word
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		err     error
	}{
		{
			name: "Should return error when word length is incorrect",
			fields: fields{
				Word: "hello",
			},
			args: args{
				guess: "hell",
			},
			wantErr: true,
			err:     LengthIsInCorrectErr,
		},
		{
			name: "Should return error when word length is incorrect with spaces",
			fields: fields{
				Word: "hello",
			},
			args: args{
				guess: "hell ",
			},
			wantErr: true,
			err:     LengthIsInCorrectErr,
		},
		{
			name: "Should return error when already guessed correctly",
			fields: fields{
				Word: "test",
				WordGuesses: []WordGuess{{
					Letters: []Letter{
						NewLetter('t', Correct),
						NewLetter('e', Correct),
						NewLetter('s', Correct),
						NewLetter('t', Correct),
					},
					Guess: "test",
				}},
			},
			args: args{
				guess: "test",
			},
			wantErr: true,
			err:     AlreadyGuessedCorrectlyErr,
		},
		{
			name: "Should return error when max word guesses exceeded",
			fields: fields{
				Word:           "t",
				MaxWordGuesses: 1,
				WordGuesses: []WordGuess{
					{
						Letters: []Letter{
							NewLetter('a', Absent),
						},
						Guess: "a",
					},
				},
			},
			args: args{
				guess: "t",
			},
			wantErr: true,
			err:     MaxWordGuessesExceededErr,
		},
		{
			name: "Should return no error",
			fields: fields{
				Word:           "hello",
				MaxWordGuesses: 6,
				WordGuesses:    []WordGuess{},
			},
			args: args{
				guess: "world",
			},
			wantErr: false,
			err:     nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := Game{
				ID:             tt.fields.Id,
				Word:           tt.fields.Word,
				WordGuesses:    tt.fields.WordGuesses,
				MaxWordGuesses: tt.fields.MaxWordGuesses,
			}
			err := g.MakeGuess(tt.args.guess)
			assert.Equal(t, tt.wantErr, err != nil)
			if tt.wantErr {
				assert.ErrorIs(t, err, tt.err)
			}
		})
	}
}
