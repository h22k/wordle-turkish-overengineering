import React, { createContext, useState } from 'react'
import { MAX_ATTEMPTS, WORD_LENGTH } from '../../gameConfig'
import { KeyboardContextType, LetterProps, LetterStatus } from '@/types/game'

export const KeyboardContext = createContext<KeyboardContextType | undefined>(undefined)

export const KeyboardProvider: React.FC<React.PropsWithChildren<{}>> = ({ children }) => {
  const [ letters, setLetters ] = useState<LetterProps[][]>(
    Array.from({ length: MAX_ATTEMPTS }).map(() =>
      Array.from({ length: WORD_LENGTH }).map(() => ( { char: '', status: LetterStatus.EMPTY } )),
    ),
  )
  const [ currentRow, setCurrentRow ] = useState(0)

  const handleClick = (value: string) => {
    const updatedLetters = [ ...letters ]
    const firstEmptyIndex = updatedLetters[currentRow].findIndex((letter) => letter.char === '')

    if ( value === 'BACKSPACE' ) {
      if ( firstEmptyIndex === -1 ) {
        updatedLetters[currentRow][WORD_LENGTH - 1].char = ''
      }
      else {
        updatedLetters[currentRow][firstEmptyIndex - 1].char = ''
      }
      setLetters(updatedLetters)
    }
    else if ( value === 'ENTER' ) {
      const currentRowLetters = updatedLetters[currentRow]
      const filledLetters = currentRowLetters.filter(letter => letter.char !== '')

      if ( filledLetters.length === WORD_LENGTH ) {
        moveToNextRow()
      }
      else {
        // toast.error('Not enough letters', {
        //   position: 'top-center',
        //   autoClose: 3000,
        //   hideProgressBar: true,
        //   closeOnClick: true,
        // })

        document.getElementById(`row-${ currentRow }`)?.classList.add('animate-shake')

        setTimeout(() => {
          document.getElementById(`row-${ currentRow }`)?.classList.remove('animate-shake')
        }, 500)
      }
    }
    else {
      if ( firstEmptyIndex !== -1 ) {
        updatedLetters[currentRow][firstEmptyIndex].char = value
        setLetters(updatedLetters)
      }
    }
  }

  const handleChange = (rowIndex: number, index: number, value: string) => {
    const updatedLetters = [ ...letters ]
    updatedLetters[rowIndex][index].char = value
    setLetters(updatedLetters)
  }

  const moveToNextRow = () => {
    if ( currentRow < MAX_ATTEMPTS - 1 ) {
      setCurrentRow(currentRow + 1)
    }
  }

  return (
    <KeyboardContext.Provider value={ { letters, handleClick, handleChange, currentRow, moveToNextRow } }>
      { children }
    </KeyboardContext.Provider>
  )
}
