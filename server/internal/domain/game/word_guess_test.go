package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewWordGuess(t *testing.T) {
	type args struct {
		word  Word
		guess Word
	}
	tests := []struct {
		name               string
		args               args
		letterCount        int
		isCorrect          bool
		correctLetterCount int
		presentLetterCount int
		absentLetterCount  int
	}{
		{
			name: "Should return correct guess",
			args: args{
				word:  "hello",
				guess: "hello",
			},
			letterCount:        5,
			isCorrect:          true,
			correctLetterCount: 5,
		},
		{
			name: "Should return correct number of present and absent letters",
			args: args{
				word:  "hello",
				guess: "hallo",
			},
			isCorrect:          false,
			letterCount:        5,
			correctLetterCount: 4,
			presentLetterCount: 0,
			absentLetterCount:  1,
		},
		{
			name: "Should return correct number of present and absent letters with duplicates",
			args: args{
				word:  "hello",
				guess: "olleh",
			},
			isCorrect:          false,
			letterCount:        5,
			correctLetterCount: 1,
			presentLetterCount: 4,
			absentLetterCount:  0,
		},
		{
			name: "Should return correct number of present with duplicates",
			args: args{
				word:  "aabbcde",
				guess: "bbxxbxx",
			},
			isCorrect:          false,
			letterCount:        7,
			correctLetterCount: 0,
			presentLetterCount: 2,
			absentLetterCount:  5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			wordGuess := NewWordGuess(tt.args.word, tt.args.guess)
			assert.Equal(t, tt.letterCount, len(wordGuess.Letters))

			assert.Equal(t, tt.isCorrect, wordGuess.IsCorrect())

			assert.Equal(t, tt.correctLetterCount, countLetters(wordGuess.Letters, Correct))
			assert.Equal(t, tt.presentLetterCount, countLetters(wordGuess.Letters, Present))
			assert.Equal(t, tt.absentLetterCount, countLetters(wordGuess.Letters, Absent))
		})
	}
}

func countLetters(letters []Letter, status LetterStatus) int {
	count := 0
	for _, letter := range letters {
		if letter.Status == status {
			count++
		}
	}
	return count
}

func Test_letterFrequencies(t *testing.T) {
	type args struct {
		word Word
	}
	tests := []struct {
		name string
		args args
		want map[rune]int
	}{
		{
			name: "Should return letter frequencies",
			args: args{
				word: "hello",
			},
			want: map[rune]int{
				'h': 1,
				'e': 1,
				'l': 2,
				'o': 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equalf(t, tt.want, letterFrequencies(tt.args.word), "letterFrequencies(%v)", tt.args.word)
		})
	}
}
