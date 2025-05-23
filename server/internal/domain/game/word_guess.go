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
	correctLetterIndices := make(map[int]struct{})

	runeIndex := 0

	for _, char := range guess {
		if runeWord[runeIndex] == char {
			letters[runeIndex] = NewLetter(char, Correct)
			correctLetterIndices[runeIndex] = struct{}{}
			letterMap[char]--
		}
		runeIndex++
	}

	runeIndex = 0 // reset index for the second check

	for _, char := range guess {
		if _, ok := correctLetterIndices[runeIndex]; ok {
			runeIndex++
			continue
		}
		status := Absent

		if letterMap[char] > 0 {
			status = Present
			letterMap[char]--
		}

		letters[runeIndex] = NewLetter(char, status)
		runeIndex++
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
