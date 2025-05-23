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

func NewWordGuess(word, guess Word) WordGuess {
	guessRunes := []rune(guess.String())
	wordRunes := []rune(word.String())

	letters := make([]Letter, len(guessRunes))
	letterMap := letterFrequencies(word)

	for i, char := range guessRunes {
		if i < len(wordRunes) && wordRunes[i] == char {
			letters[i] = NewLetter(char, Correct)
			letterMap[char]--
		}
	}

	for i, char := range guessRunes {
		if letters[i].Status != "" {
			continue
		}

		status := Absent

		if letterMap[char] > 0 {
			status = Present
			letterMap[char]--
		}

		letters[i] = NewLetter(char, status)
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

	for _, letter := range word {
		letterMap[letter]++
	}

	return letterMap
}
