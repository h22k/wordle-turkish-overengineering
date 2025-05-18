import LetterBox from './LetterBox'
import { LetterProps, LetterRowProps } from '../../types/game'
import { useKeyboard } from '../../hooks/useKeyboard'

function LetterRow({ rowIndex }: LetterRowProps) {
  const { letters, activeBoxIndex, currentRow, shakeRowIndex } = useKeyboard()
  const shouldShake = shakeRowIndex === rowIndex

  return (
    <div className={ `flex gap-[5px] ${ shouldShake ? 'animate-wiggle' : '' }` } id={ `row-${ rowIndex }` }>
      { letters[rowIndex].map((letter: LetterProps, index: number) => (
        <LetterBox
          key={ index }
          letter={ letter.char }
          status={ letter.status }
          index={ index }
          rowIndex={ rowIndex }
          isFirstBox={ rowIndex === currentRow && index === activeBoxIndex }
        />
      )) }
    </div>
  )
}

export default LetterRow
