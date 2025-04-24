import LetterBox from './LetterBox'
import { useKeyboard } from '../../context/keyboardContext'
import React from 'react'

type LetterStatus = 'empty' | 'correct' | 'present' | 'absent'

interface LetterRowProps {
  rowIndex: number
  isFirstRow: boolean
  currentRow: number
}

interface LetterProps {
  status: LetterStatus,
  char: string
}

function LetterRow({ rowIndex, isFirstRow }: LetterRowProps) {
  const { letters, handleChange, currentRow } = useKeyboard()

  const isCurrentRow = rowIndex === currentRow

  return (
    <div className="flex gap-[5px]">
      { letters[rowIndex].map((letter: LetterProps, index: number) => (
        <LetterBox
          key={ index }
          letter={ letter.char }
          status={ letter.status }
          index={ index }
          isFirstBox={ isFirstRow && index === 0 }
          onChange={ isCurrentRow ? (value) => handleChange(rowIndex, index, value) : undefined }
        />
      )) }
    </div>
  )
}

export default LetterRow
