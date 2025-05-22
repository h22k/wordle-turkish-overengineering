import React, { createContext, useEffect, useState } from 'react'
import { MAX_ATTEMPTS, VALID_LETTERS_REGEX, WORD_LENGTH } from '../gameConfig'
import { APIGuess, KeyboardContextType, LetterProps, LetterStatus } from '../types/game'
import { useToast } from '../hooks/useToast'
import axiosClient from '../api/axiosClient'
import { getUpdatedAlphabet } from '../utils/updateAlphabet'

export const KeyboardContext = createContext<KeyboardContextType | undefined>(undefined)

export const KeyboardProvider: React.FC<React.PropsWithChildren<object>> = ({ children }) => {
  const [maxAttempts, setMaxAttempts] = useState(MAX_ATTEMPTS)
  const [wordLength, setWordLength] = useState(WORD_LENGTH)
  const [loading, setLoading] = useState(false)
  const [alphabet, setAlphabet] = useState<Record<string, LetterStatus>>({})
  const [letters, setLetters] = useState<LetterProps[][]>(
    Array.from({ length: maxAttempts }).map(() =>
      Array.from({ length: wordLength }).map(() => ({ char: '', status: LetterStatus.EMPTY }))
    )
  )
  const [currentRow, setCurrentRow] = useState(0)
  const [activeBoxIndex, setActiveBoxIndex] = useState(0)
  const [shakeRowIndex, setShakeRowIndex] = useState<number | null>(null)
  const [submittedRow, setSubmittedRow] = useState<number | null>(null)
  const { notify } = useToast()

  useEffect(() => {
    const fetchGame = async () => {
      try {
        setLoading(true)
        const res = await axiosClient.get('/game/game')

        const { guesses, max_guesses }: { guesses: APIGuess[]; max_guesses: number } = res.data.data
        setMaxAttempts(max_guesses)
        setWordLength(max_guesses - 1)
        const letters = guesses.map((guess) => guess.letters)

        const remaining = Array.from({ length: max_guesses - letters.length }).map(() =>
          Array.from({ length: max_guesses - 1 }).map(() => ({
            char: '',
            status: LetterStatus.EMPTY,
          }))
        )

        setLetters([...letters, ...remaining])
        setCurrentRow(letters.length)

        const newAlphabet = getUpdatedAlphabet({}, letters)
        setAlphabet(newAlphabet)
      } catch (err: any) {
        const message = err?.message || 'Oyun Yüklenemedi.'
        notify(message)
      } finally {
        setLoading(false)
      }
    }

    fetchGame()
  }, [])

  const triggerShake = (row: number) => {
    setShakeRowIndex(row)
    setTimeout(() => setShakeRowIndex(null), 500)
  }

  const focusInput = (row: number, index: number) => {
    setTimeout(() => {
      const input = document.getElementById(`input-${row}-${index}`) as HTMLInputElement | null
      input?.focus()
    }, 0)
  }

  const moveToNextRow = () => {
    if (currentRow < maxAttempts - 1) {
      setCurrentRow(currentRow + 1)
      setActiveBoxIndex(0)
      focusInput(currentRow + 1, 0)
    }
  }

  const addLetter = (value: string) => {
    if (activeBoxIndex > wordLength) return

    const updatedLetters = [...letters]
    const currentBox = updatedLetters[currentRow]?.[activeBoxIndex]

    if (currentBox?.char !== '') return

    if (currentBox) {
      currentBox.char = value
    }
    setLetters(updatedLetters)

    const isLastBox = activeBoxIndex === wordLength - 1
    if (!isLastBox) {
      setActiveBoxIndex(activeBoxIndex + 1)
      focusInput(currentRow, activeBoxIndex + 1)
    }
  }

  const deleteLetter = () => {
    if (activeBoxIndex < 0 || activeBoxIndex >= wordLength || currentRow >= maxAttempts) return

    const updatedLetters = [...letters]
    const currentBox = updatedLetters[currentRow]?.[activeBoxIndex]

    if (currentBox?.char) {
      currentBox.char = ''
      setLetters(updatedLetters)
    } else if (activeBoxIndex > 0) {
      const previousBox = updatedLetters[currentRow]?.[activeBoxIndex - 1]
      if (previousBox) {
        previousBox.char = ''
        const newIndex = activeBoxIndex - 1
        setLetters(updatedLetters)
        setActiveBoxIndex((prev) => prev - 1)
        focusInput(currentRow, newIndex)
      }
    }
  }

  const submitWord = async () => {
    const currentWord = letters[currentRow]
    const filledLetters = currentWord.filter((letter) => letter.char !== '')

    if (filledLetters.length !== wordLength) {
      triggerShake(currentRow)
      return
    }

    try {
      const res = await axiosClient.post('/game/guess', {
        guess: currentWord?.map((l) => l.char).join(''),
      })

      const updated = [...letters]
      const newLetters = res?.data?.data?.letters
      updated[currentRow] = newLetters
      setLetters(updated)

      const newAlphabet = getUpdatedAlphabet(alphabet, [newLetters])
      setAlphabet(newAlphabet)

      setSubmittedRow(currentRow)
      moveToNextRow()
    } catch (err: any) {
      const message = err?.response?.data?.error || err?.message || 'Bir hata oluştu'

      notify(message)
      triggerShake(currentRow)
    }
  }

  const handleChange = async (value: string) => {
    const upperValue = value.toLocaleUpperCase('tr')

    if (upperValue === 'BACKSPACE') {
      deleteLetter()
      return
    } else if (upperValue === 'ENTER') {
      await submitWord()
    } else if (VALID_LETTERS_REGEX.test(upperValue)) {
      addLetter(upperValue)
      setActiveBoxIndex(Math.min(activeBoxIndex + 1, wordLength - 1))
    }
  }

  const getFirstEmptyBoxIndex = (row: number) => {
    return letters?.[row]?.findIndex((letter) => letter.char === '')
  }

  return (
    <KeyboardContext.Provider
      value={{
        letters,
        handleChange,
        currentRow,
        activeBoxIndex,
        setActiveBoxIndex,
        shakeRowIndex,
        triggerShake,
        getFirstEmptyBoxIndex,
        submittedRow,
        maxAttempts,
        wordLength,
        loading,
        alphabet,
      }}
    >
      {children}
    </KeyboardContext.Provider>
  )
}
