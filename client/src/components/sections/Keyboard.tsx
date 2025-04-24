import Key from '../common/Key'
import { KEYBOARD_LAYOUT } from '../../../gameConfig'
import { useKeyboard } from '../../context/keyboardContext'
import { useEffect } from 'react'

function Keyboard() {
  const { handleClick } = useKeyboard()

  useEffect(() => {
    const handleKeyDown = (e: KeyboardEvent) => {
      const key = e.key.toUpperCase()

      if (['BACKSPACE', 'ENTER', ...KEYBOARD_LAYOUT?.flat()]?.includes(key)) {
        handleClick(key)
      }
    }

    window.addEventListener('keydown', handleKeyDown)
    return () => window.removeEventListener('keydown', handleKeyDown)
  }, [ handleClick ])

  return (
    <div className="flex flex-col gap-2 items-center">
      { KEYBOARD_LAYOUT.map((row, rowIndex) => (
        <div key={ rowIndex } className="flex gap-1.5">
          { row.map((key) => (
            <Key key={ key } value={ key } onClick={ handleClick }/>
          )) }
        </div>
      )) }
    </div>
  )
}

export default Keyboard
