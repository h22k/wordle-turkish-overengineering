import LetterBox from './LetterBox'
import { LetterProps, LetterRowProps } from '../../types/game'
import { useKeyboard } from '../../hooks/useKeyboard'

function LetterRow({ rowIndex, isFirstRow }: LetterRowProps) {
  const { letters } = useKeyboard()

  return (
    <div className="flex gap-[5px]" id={ `row-${ rowIndex }` }>
      { letters[rowIndex].map((letter: LetterProps, index: number) => (
        <LetterBox
          key={ index }
          letter={ letter.char }
          status={ letter.status }
          index={ index }
          isFirstBox={ isFirstRow && index === 0 }
        />
      )) }
    </div>
  )
}

export default LetterRow
