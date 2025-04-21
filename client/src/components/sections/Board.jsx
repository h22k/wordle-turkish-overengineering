import LetterRow from '../common/LetterRow.jsx'
import { MAX_ATTEMPTS } from '../../../gameConfig.js'

function Board() {
  return (
    <div className="grid gap-2">
      { Array.from({ length: MAX_ATTEMPTS }).map((_, i) => (
        <LetterRow key={ i } isFirstRow={ i === 0 }/>
      )) }
    </div>
  )
}

export default Board
