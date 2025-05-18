import { useEffect } from 'react'
import Key from '../common/Key'
import { KEYBOARD_LAYOUT } from '../../gameConfig'
import { useKeyboardEvents } from '../../hooks/useKeyboardEvents'

function Keyboard() {
  const { processKey } = useKeyboardEvents()

  useEffect(() => {
    const handleKeyDown = (e: KeyboardEvent) => {
      if ( document.activeElement?.tagName === 'INPUT' ) return
      processKey(e.key)
    }

    window.addEventListener('keydown', handleKeyDown)
    return () => window.removeEventListener('keydown', handleKeyDown)
  }, [ processKey ])

  return (
    <div className="flex flex-col gap-2 items-center w-full sm:w-auto">
      { KEYBOARD_LAYOUT.map((row, rowIndex) => (
        <div key={ rowIndex } className="flex gap-1.5 w-full sm:w-auto">
          { row.map((key) => (
            <Key key={ key } value={ key } onClick={ processKey }/>
          )) }
        </div>
      )) }
    </div>
  )
}

export default Keyboard
