import { useKeyboard } from './useKeyboard'
import React, { useState } from 'react'
import { VALID_LETTERS_REGEX } from '../gameConfig'

export const useKeyboardEvents = (inputRef?: React.RefObject<HTMLInputElement | null>) => {
  const { handleChange } = useKeyboard()

  const processKey = (key: string) => {
    const upperKey = key.toUpperCase()
    if ( VALID_LETTERS_REGEX.test(upperKey) ) {
      handleChange(upperKey)
      setTimeout(() => {
        ( inputRef?.current?.nextElementSibling as HTMLInputElement | null )?.focus()
      }, 0)
    }

    else if ( upperKey === 'BACKSPACE' ) {
      handleChange('BACKSPACE')
      setTimeout(() => {
        ( inputRef?.current?.previousElementSibling as HTMLInputElement | null )?.focus()
      }, 0)
    }

    else if ( upperKey === 'ENTER' ) {
      handleChange('ENTER')
    }
  }

  return { processKey }
}
