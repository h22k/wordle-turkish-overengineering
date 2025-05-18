import React, { useEffect, useRef, useState } from 'react'
import { STATUS_COLOR } from '../../gameConfig'
import { LetterBoxProps } from '../../types/game'
import { useKeyboardEvents } from '../../hooks/useKeyboardEvents'
import { useKeyboard } from '../../hooks/useKeyboard'

function LetterBox({ letter, status, isFirstBox, index, rowIndex }: LetterBoxProps) {
  const inputRef = useRef<HTMLInputElement>(null)
  const isSpecialKeyRef = useRef<'BACKSPACE' | 'ENTER' | null>(null)
  const { processKey } = useKeyboardEvents(inputRef)
  const { setActiveBoxIndex } = useKeyboard()
  const [ isAnimating, setIsAnimating ] = useState(false)

  const handleFocus = () => {
    if ( typeof index === 'number' ) {
      setActiveBoxIndex(index)
    }
  }

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    processKey(e.target.value.toUpperCase())
  }

  const handleKeyDown = (e: React.KeyboardEvent<HTMLInputElement>) => {
    const key = e.key.toUpperCase()

    if ( key === 'ENTER' || key === 'BACKSPACE' ) {
      processKey(key)
      return
    }
    else {
      handleAnimation()
      isSpecialKeyRef.current = null
    }
  }

  const handleAnimation = () => {
    if ( !inputRef?.current?.value ) {
      setIsAnimating(true)
      setTimeout(() => {
        setIsAnimating(false)
      }, 50)
    }
  }

  return (
    <input
      autoFocus={ isFirstBox }
      ref={ inputRef }
      type="text"
      id={ `input-${ rowIndex }-${ index }` }
      maxLength={ 1 }
      value={ letter }
      onChange={ handleChange }
      onKeyDown={ handleKeyDown }
      onMouseDown={ (e: React.MouseEvent<HTMLInputElement>) => e.preventDefault() }
      onFocus={ handleFocus }
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
