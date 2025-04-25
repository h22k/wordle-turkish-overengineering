import React, { useRef, useEffect, useState } from 'react'
import { STATUS_COLOR } from '../../gameConfig'
import { LetterBoxProps } from '../../types/game'
import { useKeyboardEvents } from '../../hooks/useKeyboardEvents'

function LetterBox({ letter, status, isFirstBox }: LetterBoxProps) {
  const inputRef = useRef<HTMLInputElement>(null)
  const [ isAnimating, setIsAnimating ] = useState(false)
  const { processKey } = useKeyboardEvents()

  useEffect(() => {
    if ( isFirstBox ) {
      inputRef.current?.focus()
    }
  }, [ isFirstBox ])


  const handleKeyDown = (e: React.KeyboardEvent<HTMLInputElement>) => {
    processKey(e.key)
    setIsAnimating(true)
    setTimeout(() => {
      ( inputRef.current?.nextElementSibling as HTMLInputElement | null )?.focus()
      setIsAnimating(false)
    }, 50)
  }

  return (
    <input
      ref={ inputRef }
      type="text"
      maxLength={ 1 }
      value={ letter }
      onKeyDown={ handleKeyDown }
      className={ `
        w-[52px] h-[52px] text-center uppercase font-bold text-[2rem] text-white
        focus:outline-none transition-transform duration-75
        caret-transparent cursor-default select-none ${ STATUS_COLOR[status] }
        ${ isAnimating ? 'scale-110' : 'scale-100' }
      ` }
    />
  )
}

export default LetterBox
