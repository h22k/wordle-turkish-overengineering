import React, { useEffect, useRef, useState } from 'react'
import { STATUS_COLOR } from '../../gameConfig'
import { LetterBoxProps } from '../../types/game'
import { useKeyboardEvents } from '../../hooks/useKeyboardEvents'
import { useKeyboard } from '../../hooks/useKeyboard'

function LetterBox({ letter, status, isFirstBox, index, rowIndex }: LetterBoxProps) {
  const inputRef = useRef<HTMLInputElement>(null)
  const isSpecialKeyRef = useRef<'BACKSPACE' | 'ENTER' | null>(null)
  const { processKey } = useKeyboardEvents(inputRef)
  const { setActiveBoxIndex, submittedRow } = useKeyboard()
  const [isAnimating, setIsAnimating] = useState(false)
  const [isFlipping, setIsFlipping] = useState(false)
  const prevLetterRef = useRef<string>('')
  const flipDelay = 150

  useEffect(() => {
    const shouldFlip = rowIndex === submittedRow && letter && prevLetterRef.current !== letter
    if (shouldFlip) {
      const delay = setTimeout(() => setIsFlipping(true), (index || 0) * 150)
      const clear = setTimeout(() => setIsFlipping(false), 600 + (index || 0) * 150)
      prevLetterRef.current = letter
      return () => {
        clearTimeout(delay)
        clearTimeout(clear)
      }
    }
  }, [letter, submittedRow, rowIndex, index])

  const handleFocus = () => {
    if (typeof index === 'number') {
      setActiveBoxIndex(index)
    }
  }

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    processKey(e.target.value.toLocaleUpperCase('tr'))
  }

  const handleKeyDown = (e: React.KeyboardEvent<HTMLInputElement>) => {
    const key = e.key.toLocaleUpperCase('tr')

    if (key === 'ENTER' || key === 'BACKSPACE') {
      processKey(key)
      return
    } else {
      handleAnimation()
      isSpecialKeyRef.current = null
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
      onChange={handleChange}
      onKeyDown={handleKeyDown}
      onMouseDown={(e: React.MouseEvent<HTMLInputElement>) => e.preventDefault()}
      onFocus={handleFocus}
      className={`
        w-[52px] h-[52px] text-center font-bold text-[2rem] text-white
        focus:outline-none transition-transform duration-75
        caret-transparent cursor-default select-none ${STATUS_COLOR[status]}
        ${isAnimating ? 'scale-110' : 'scale-100'}
        ${isFlipping ? 'animate-flip' : ''}
      `}
      style={{ transitionDelay: `${flipDelay}ms` }}
    />
  )
}

export default LetterBox
