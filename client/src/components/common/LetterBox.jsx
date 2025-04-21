import { useEffect, useRef, useState } from 'react'

function LetterBox({ letter, status, index, onChange, isFirstBox }) {
  const inputRef = useRef(null)
  const [ isAnimating, setIsAnimating ] = useState(false)

  const statusColor = {
    correct: 'bg-correct',
    present: 'bg-present',
    absent: 'bg-absent',
    empty: 'bg-transparent border-2 border-absent',
  }

  useEffect(() => {
    if (isFirstBox) {
      inputRef.current?.focus()
    }
  }, [isFirstBox])

  const handleChange = (e) => {
    const value = e.target.value.toUpperCase()
    if ( /^[A-Z]$/.test(value) ) {
      onChange(index, value)
      setIsAnimating(true)
      setTimeout(() => setIsAnimating(false), 75)
      inputRef.current.nextElementSibling?.focus()
    }
    else if ( value === '' ) {
      onChange(index, '')
    }
  }

  const handleKeyDown = (e) => {
    if ( e.key === 'Backspace' && !letter ) {
      inputRef.current.previousElementSibling?.focus()
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
        caret-transparent cursor-default select-none ${ statusColor[status] } 
        ${ isAnimating ? 'scale-110' : 'scale-100' }` }
    />
  )
}

export default LetterBox
