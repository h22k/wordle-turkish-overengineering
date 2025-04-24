import { useRef, useEffect, useState } from 'react'
import { STATUS_COLOR } from '../../../gameConfig.js'

function LetterBox({ letter, status, onChange, isFirstBox }) {
  const inputRef = useRef(null)
  const [ isAnimating, setIsAnimating ] = useState(false)

  useEffect(() => {
    if ( isFirstBox ) {
      inputRef.current?.focus()
    }
  }, [ isFirstBox ])

  const handleChange = (e) => {
    const value = e.target.value.toUpperCase()
    if ( value && /^[A-Z]$/.test(value) ) {
      onChange(value)
      setIsAnimating(true)
      setTimeout(() => setIsAnimating(false), 75)
    }
  }

  const handleKeyDown = (e) => {
    if ( e.key === 'Backspace' && !letter ) {
      inputRef.current.previousElementSibling?.focus()
    }
    if ( e.key === 'Enter' ) {
      e.preventDefault()
    }
  }

  return (
    <input
      ref={ inputRef }
      type="text"
      maxLength={ 1 }
      value={ letter }
      onChange={ handleChange }
      onKeyDown={ handleKeyDown }
      className={ `w-[52px] h-[52px] text-center uppercase font-bold text-[2rem] text-white 
        focus:outline-none transition-transform duration-75 
        caret-transparent cursor-default select-none ${ STATUS_COLOR[status] } 
        ${ isAnimating ? 'scale-110' : 'scale-100' }` }
    />
  )
}

export default LetterBox
