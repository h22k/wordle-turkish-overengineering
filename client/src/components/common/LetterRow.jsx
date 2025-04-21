import { useState } from 'react'
import LetterBox from './LetterBox.jsx'
import { WORD_LENGTH } from '../../../gameConfig.js'

function LetterRow({ isFirstRow }) {
  const [ letters, setLetters ] = useState(
    Array.from({ length: WORD_LENGTH }).map(() => ( { char: '', status: 'empty' } )),
  )

  const handleChange = (index, value) => {
    const updated = [ ...letters ]
    updated[index].char = value
    setLetters(updated)
  }

  return (
    <div className="flex gap-1.5">
      { letters.map((letter, index) => (
        <LetterBox
          key={ index }
          letter={ letter.char }
          status={ letter.status }
          index={ index }
          isFirstBox={ isFirstRow && index === 0 }
          onChange={ handleChange }
        />
      )) }
    </div>
  )
}

export default LetterRow
