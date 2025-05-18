import React, { useEffect, useRef } from 'react'
import { STATUS_COLOR } from '../../gameConfig'
import { LetterBoxProps } from '../../types/game'
import { useKeyboardEvents } from '../../hooks/useKeyboardEvents'
import { useKeyboard } from '../../hooks/useKeyboard'

function LetterBox({ letter, status, isFirstBox, index, rowIndex }: LetterBoxProps) {
  const inputRef = useRef<HTMLInputElement>(null)
  const { processKey, isAnimating } = useKeyboardEvents(inputRef)
  const { setActiveBoxIndex } = useKeyboard()

  useEffect(() => {
    if (isFirstBox) {
      inputRef.current?.focus()
    }
  }, [isFirstBox])

  const handleFocus = () => {
    if (typeof index === 'number') {
      setActiveBoxIndex(index)
    }
  }

  const handleKeyDown = (e: React.KeyboardEvent<HTMLInputElement>) => {
    processKey(e.key)
  }

  return (
    <input
      ref={inputRef}
      type="text"
      id={`input-${rowIndex}-${index}`}
      maxLength={1}
      value={letter}
      onKeyDown={handleKeyDown}
      onMouseDown={(e: React.MouseEvent<HTMLInputElement>) => e.preventDefault()}
      onFocus={handleFocus}
      className={`
        w-[52px] h-[52px] text-center uppercase font-bold text-[2rem] text-white
        focus:outline-none transition-transform duration-75
        caret-transparent cursor-default select-none ${STATUS_COLOR[status]}
        ${isAnimating ? 'scale-110' : 'scale-100'}
      `}
    />
  )
}

export default LetterBox
