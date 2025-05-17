import React, { createContext, useState } from 'react'
import { MAX_ATTEMPTS, VALID_LETTERS_REGEX, WORD_LENGTH } from '../gameConfig'
import { KeyboardContextType, LetterProps, LetterStatus } from '../types/game'
import { useToast } from '../hooks/useToast'

export const KeyboardContext = createContext<KeyboardContextType | undefined>(undefined)

export const KeyboardProvider: React.FC<React.PropsWithChildren<{}>> = ({ children }) => {
  const [ letters, setLetters ] = useState<LetterProps[][]>(
    Array.from({ length: MAX_ATTEMPTS }).map(() =>
      Array.from({ length: WORD_LENGTH }).map(() => ( { char: '', status: LetterStatus.EMPTY } )),
    ),
  )
  const [ currentRow, setCurrentRow ] = useState(0)
  const [ activeBoxIndex, setActiveBoxIndex ] = useState(0)
  const [ shakeRowIndex, setShakeRowIndex ] = useState<number | null>(null)

  const { notify } = useToast()

  const triggerShake = (row: number) => {
    setShakeRowIndex(row)
    setTimeout(() => setShakeRowIndex(null), 500)
  }

  const moveToNextRow = () => {
    if ( currentRow < MAX_ATTEMPTS - 1 ) {
      setCurrentRow(currentRow + 1)
      setActiveBoxIndex(0)
      focusInput(currentRow + 1, 0)
    }
  }

  const focusInput = (row: number, index: number) => {
    setTimeout(() => {
      const input = document.getElementById(`input-${ row }-${ index }`) as HTMLInputElement | null
      input?.focus()
    }, 0)
  }

  const addLetter = (value: string) => {
    if ( activeBoxIndex >= WORD_LENGTH ) return

    const updatedLetters = [ ...letters ]
    const currentBox = updatedLetters[currentRow][activeBoxIndex]

    if ( currentBox.char !== '' ) return

    currentBox.char = value
    setLetters(updatedLetters)

    const isLastBox = activeBoxIndex === WORD_LENGTH - 1
    if ( !isLastBox ) {
      setActiveBoxIndex(activeBoxIndex + 1)
      focusInput(currentRow, activeBoxIndex + 1)
    }
  }

  const deleteLetter = () => {
    if ( activeBoxIndex < 0 ) return

    const updatedLetters = [ ...letters ]
    const currentBox = updatedLetters[currentRow][activeBoxIndex]

    if ( currentBox?.char ) {
      currentBox.char = ''
    }
    else {
      updatedLetters[currentRow][activeBoxIndex - 1].char = ''
      setActiveBoxIndex(activeBoxIndex - 1)
      focusInput(currentRow, activeBoxIndex - 1)
    }

    setLetters(updatedLetters)
  }

  const submitWord = () => {
    const currentWord = letters[currentRow]
    const filledLetters = currentWord.filter(letter => letter.char !== '')

    if ( filledLetters.length === WORD_LENGTH ) {
      moveToNextRow()
    }
    else {
      notify('Harf sayısı yetersiz')
      triggerShake(currentRow)
    }
  }

  const handleChange = (value: string) => {
    const upperValue = value.toUpperCase()

    if ( upperValue === 'BACKSPACE' ) {
      deleteLetter()
    }
    else if ( upperValue === 'ENTER' ) {
      submitWord()
    }
    else if ( VALID_LETTERS_REGEX.test(upperValue) ) {
      addLetter(upperValue)
      setActiveBoxIndex(Math.min(activeBoxIndex + 1, WORD_LENGTH - 1))
    }
  }

  return (
    <KeyboardContext.Provider
      value={ {
        letters,
        handleChange,
        currentRow,
        moveToNextRow,
        activeBoxIndex,
        setActiveBoxIndex,
        shakeRowIndex,
        triggerShake,
      } }>
      { children }
    </KeyboardContext.Provider>
  )
}
