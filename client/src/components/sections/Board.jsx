import LetterRow from '../common/LetterRow.jsx'
import { MAX_ATTEMPTS } from '../../../gameConfig.js'
import { useKeyboard } from '../../context/keyboardContext.jsx'

function Board() {
  const { currentRow } = useKeyboard()

  return (
    <div className="grid gap-[5px]">
      { Array.from({ length: MAX_ATTEMPTS }).map((_, i) => (
        <LetterRow
          key={ i }
          rowIndex={ i }
          isFirstRow={ i === 0 }
          currentRow={ currentRow }
        />
      )) }
    </div>
  )
}

export default Board
