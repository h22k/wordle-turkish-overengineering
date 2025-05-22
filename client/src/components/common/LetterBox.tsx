import React, { useEffect, useRef, useState } from 'react'
import { STATUS_COLOR } from '../../gameConfig'
import { LetterBoxProps, LetterStatus } from '../../types/game'
import { useKeyboard } from '../../hooks/useKeyboard'

function LetterBox({ letter, status, isFirstBox, index, rowIndex }: LetterBoxProps) {
  const inputRef = useRef<HTMLInputElement>(null)
  const { setActiveBoxIndex, handleChange } = useKeyboard()

  const [isAnimating, setIsAnimating] = useState(false)
  const [isFlipping, setIsFlipping] = useState(false)
  const [showStatusColor, setShowStatusColor] = useState(false)
  const flipDelay = (index || 0) * 150

  useEffect(() => {
    const isFilled = letter !== ''
    const hasStatus = status !== LetterStatus.EMPTY
    const shouldFlip = isFilled && hasStatus

    if (shouldFlip) {
      const flipTimer = setTimeout(() => {
        setIsFlipping(true)
        const colorTimer = setTimeout(() => {
          setIsFlipping(false)
          setShowStatusColor(true)
        }, 500)

        return () => {
          clearTimeout(flipTimer)
          clearTimeout(colorTimer)
        }
      }, flipDelay)
    }
  }, [letter, status, rowIndex])

  const handleFocus = () => {
    if (typeof index === 'number') {
      setActiveBoxIndex(index)
    }
  }

  const handleKeyDown = (e: React.KeyboardEvent<HTMLInputElement>) => {
    const key = e.key.toLocaleUpperCase('tr')

    if (key === 'ENTER' || key === 'BACKSPACE') {
      handleChange(key)
      return
    } else {
      handleAnimation()
    }
  }

  const handleAnimation = () => {
    if (!inputRef?.current?.value) {
      setIsAnimating(true)
      setTimeout(() => {
        setIsAnimating(false)
      }, 50)
    }
  }

  return (
    <input
      autoFocus={isFirstBox}
      ref={inputRef}
      type="text"
      id={`input-${rowIndex}-${index}`}
      maxLength={1}
      value={letter.toLocaleUpperCase('tr')}
      onMouseDown={(e) => e.preventDefault()}
      onChange={(e) => handleChange(e.target.value.toLocaleUpperCase('tr'))}
      onKeyDown={handleKeyDown}
      onFocus={handleFocus}
      className={`
        w-[52px] h-[52px] text-center font-bold text-[2rem] text-white
        focus:outline-none transition-transform duration-75
        caret-transparent cursor-default select-none 
        ${isAnimating ? 'scale-110' : 'scale-100'}
        ${isFlipping ? 'animate-flip' : ''}
        ${showStatusColor ? STATUS_COLOR[status] : STATUS_COLOR[LetterStatus.EMPTY]}
        delay-[${flipDelay}ms]
      `}
    />
  )
}

export default LetterBox
