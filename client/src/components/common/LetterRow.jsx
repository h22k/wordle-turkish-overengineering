import LetterBox from './LetterBox.jsx'
import { useKeyboard } from '../../context/keyboardContext.jsx'

function LetterRow({ rowIndex, isFirstRow }) {
  const { letters, handleChange, currentRow } = useKeyboard()

  const isCurrentRow = rowIndex === currentRow

  return (
    <div className="flex gap-[5px]">
      { letters[rowIndex].map((letter, index) => (
        <LetterBox
          key={ index }
          letter={ letter.char }
          status={ letter.status }
          index={ index }
          isFirstBox={ isFirstRow && index === 0 }
          onChange={ isCurrentRow ? (value) => handleChange(rowIndex, index, value) : null }
        />
      )) }
    </div>
  )
}

export default LetterRow
