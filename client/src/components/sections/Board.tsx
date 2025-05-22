import LetterRow from '../common/LetterRow'
import { useKeyboard } from '../../hooks/useKeyboard'
import Spinner from '../common/Spinner'

function Board() {
  const { currentRow, maxAttempts, loading } = useKeyboard()

  if (loading)
    return (
      <div className="h-screen flex items-center justify-center">
        <Spinner />
      </div>
    )

  return (
    <div className="grid gap-[5px]">
      {Array.from({ length: maxAttempts }).map((_, i) => (
        <LetterRow key={i} rowIndex={i} currentRow={currentRow} />
      ))}
    </div>
  )
}

export default Board
