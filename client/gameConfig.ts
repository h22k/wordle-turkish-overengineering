export const WORD_LENGTH: number = +import.meta.env.VITE_WORD_LENGTH
export const MAX_ATTEMPTS: number = +import.meta.env.VITE_MAX_ATTEMPTS
export const KEYBOARD_LAYOUT: string[][] = [
  ['Q', 'W', 'E', 'R', 'T', 'Y', 'U', 'I', 'O', 'P'],
  ['A', 'S', 'D', 'F', 'G', 'H', 'J', 'K', 'L'],
  ['ENTER', 'Z', 'X', 'C', 'V', 'B', 'N', 'M', 'BACKSPACE'],
]
export const STATUS_COLOR: Record<string, string> = {
  correct: 'bg-correct',
  present: 'bg-present',
  absent: 'bg-absent',
  empty: 'bg-transparent border-2 border-absent',
}
