import LetterBox from './LetterBox'
import { LetterProps, LetterRowProps } from '../../types/game'
import { useKeyboard } from '../../hooks/useKeyboard'
import { useEffect, useState } from 'react'

function LetterRow({ rowIndex }: LetterRowProps) {
  const { letters, currentRow, shakeRowIndex, getFirstEmptyBoxIndex } = useKeyboard()
  const shouldShake = shakeRowIndex === rowIndex
  const [emptyBox, setEmptyBox] = useState(-1)

  useEffect(() => {
    setEmptyBox(getFirstEmptyBoxIndex(rowIndex))
  }, [letters, rowIndex])

  return (
    <div className={`flex gap-[5px] ${shouldShake ? 'animate-wiggle' : ''}`} id={`row-${rowIndex}`}>
      {letters?.[rowIndex]?.map((letter: LetterProps, index: number) => (
        <LetterBox
          key={index}
          letter={letter.char}
          status={letter.status}
          index={index}
          rowIndex={rowIndex}
          isFirstBox={rowIndex === currentRow && index === emptyBox}
        />
      ))}
    </div>
  )
}

export default LetterRow
