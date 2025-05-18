import React, { createContext, useEffect, useState } from 'react'
import { MAX_ATTEMPTS, VALID_LETTERS_REGEX, WORD_LENGTH } from '../gameConfig'
import { APIGuess, KeyboardContextType, LetterProps, LetterStatus } from '../types/game'
import { useToast } from '../hooks/useToast'
import axiosClient from '../api/axiosClient'

export const KeyboardContext = createContext<KeyboardContextType | undefined>(undefined)

export const KeyboardProvider: React.FC<React.PropsWithChildren<object>> = ({ children }) => {
  const [letters, setLetters] = useState<LetterProps[][]>(
    Array.from({ length: MAX_ATTEMPTS }).map(() =>
      Array.from({ length: WORD_LENGTH }).map(() => ({ char: '', status: LetterStatus.EMPTY }))
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
        const res = await axiosClient.get('/game/game')

        const { guesses, max_guesses }: { guesses: APIGuess[]; max_guesses: number } = res.data.data

        const filled = guesses.map((guess) =>
          guess.letters.map((letter) => ({
            char: letter.char,
            status: letter.status,
          }))
        )

        const remaining = Array.from({ length: max_guesses - filled.length }).map(() =>
          Array.from({ length: WORD_LENGTH }).map(() => ({
            char: '',
            status: LetterStatus.EMPTY,
          }))
        )

        setLetters([...filled, ...remaining])
        setCurrentRow(filled.length)
      } catch (err: any) {
        const message = err?.message || 'Oyun Yüklenemedi.'
        notify(message)
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
    if (currentRow < MAX_ATTEMPTS - 1) {
      setCurrentRow(currentRow + 1)
      setActiveBoxIndex(0)
      focusInput(currentRow + 1, 0)
    }
  }

  const addLetter = (value: string) => {
    if (activeBoxIndex >= WORD_LENGTH) return

    const updatedLetters = [...letters]
    const currentBox = updatedLetters[currentRow][activeBoxIndex]

    if (currentBox.char !== '') return

    currentBox.char = value
    setLetters(updatedLetters)

    const isLastBox = activeBoxIndex === WORD_LENGTH - 1
    if (!isLastBox) {
      setActiveBoxIndex(activeBoxIndex + 1)
      focusInput(currentRow, activeBoxIndex + 1)
    }
  }

  const deleteLetter = () => {
    if (activeBoxIndex < 0) return

    const updatedLetters = [...letters]
    const currentBox = updatedLetters[currentRow][activeBoxIndex]

    if (currentBox?.char) {
      currentBox.char = ''
    } else {
      updatedLetters[currentRow][activeBoxIndex - 1].char = ''
      setActiveBoxIndex(activeBoxIndex - 1)
      focusInput(currentRow, activeBoxIndex - 1)
    }

    setLetters(updatedLetters)
  }

  const submitWord = async () => {
    const currentWord = letters[currentRow]
    const filledLetters = currentWord.filter((letter) => letter.char !== '')

    if (filledLetters.length !== WORD_LENGTH) {
      notify('Harf sayısı yetersiz')
      triggerShake(currentRow)
      return
    }

    try {
      const res = await axiosClient.post('/game/guess', {
        guess: currentWord?.map((l) => l.char).join(''),
      })

      const updated = [...letters]
      updated[currentRow] = res?.data?.data?.letters?.map((l: any) => ({
        char: l.char,
        status: l.status,
      }))
      setLetters(updated)

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
    } else if (upperValue === 'ENTER') {
      await submitWord()
    } else if (VALID_LETTERS_REGEX.test(upperValue)) {
      addLetter(upperValue)
      setActiveBoxIndex(Math.min(activeBoxIndex + 1, WORD_LENGTH - 1))
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
        moveToNextRow,
        activeBoxIndex,
        setActiveBoxIndex,
        shakeRowIndex,
        triggerShake,
        getFirstEmptyBoxIndex,
        submittedRow,
      }}
    >
      {children}
    </KeyboardContext.Provider>
  )
}
