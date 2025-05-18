import { useKeyboard } from './useKeyboard'
import React from 'react'
import { VALID_LETTERS_REGEX } from '../gameConfig'

export const useKeyboardEvents = (inputRef?: React.RefObject<HTMLInputElement | null>) => {
  const { handleChange } = useKeyboard()

  const processKey = (key: string) => {
    const upperKey = key.toLocaleUpperCase('tr')
    if (VALID_LETTERS_REGEX.test(upperKey)) {
      handleChange(upperKey)
      ;(inputRef?.current?.nextElementSibling as HTMLInputElement | null)?.focus()
    } else if (upperKey === 'BACKSPACE') {
      handleChange('BACKSPACE')
      ;(inputRef?.current?.previousElementSibling as HTMLInputElement | null)?.focus()
    } else if (upperKey === 'ENTER') {
      handleChange('ENTER')
    }
  }

  return { processKey }
}
