import LetterRow from '../common/LetterRow'
import { MAX_ATTEMPTS } from '../../gameConfig'
import { useKeyboard } from '../../hooks/useKeyboard'

function Board() {
  const { currentRow } = useKeyboard()

  return (
    <div className="grid gap-[5px]">
      {Array.from({ length: MAX_ATTEMPTS }).map((_, i) => (
        <LetterRow key={i} rowIndex={i} currentRow={currentRow} />
      ))}
    </div>
  )
}

export default Board
