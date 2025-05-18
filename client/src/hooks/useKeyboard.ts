import { useContext } from 'react'
import { KeyboardContext } from '../context/keyboardContext'

export const useKeyboard = () => {
  const context = useContext(KeyboardContext)
  if (!context) {
    throw new Error('useKeyboard must be used within a KeyboardProvider')
  }
  return context
}
