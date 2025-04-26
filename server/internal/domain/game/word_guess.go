package domain

type LetterStatus string

const (
	Correct LetterStatus = "correct"
	Present LetterStatus = "present"
	Absent  LetterStatus = "absent"
)

type Letter struct {
	Char   rune
	Status LetterStatus
}

type WordGuess struct {
	Letters []Letter
	Guess   Word
}

func NewLetter(char rune, status LetterStatus) Letter {
	return Letter{
		Char:   char,
		Status: status,
	}
}

// NewWordGuess TODO: refactor this function to make it more readable
func NewWordGuess(word, guess Word) WordGuess {
	letters := make([]Letter, guess.Len())
	letterMap := letterFrequencies(word)
	runeWord := []rune(word.String())

	for i, char := range guess.String() {
		switch true {
		case runeWord[i] == char:
			letters[i] = NewLetter(char, Correct)
			break
		case letterMap[char] > 0:
			letters[i] = NewLetter(char, Present)
			letterMap[char]--
			break
		default:
			letters[i] = NewLetter(char, Absent)
		}
	}

	return WordGuess{
		Letters: letters,
		Guess:   guess,
	}
}

func (g *WordGuess) IsCorrect() bool {
	for _, letter := range g.Letters {
		if letter.Status != Correct {
			return false
		}
	}
	return true
}

func letterFrequencies(word Word) map[rune]int {
	letterMap := make(map[rune]int)
	for _, letter := range word.String() {
		if _, ok := letterMap[letter]; ok {
			letterMap[letter]++
		} else {
			letterMap[letter] = 1
		}
	}
	return letterMap
}
