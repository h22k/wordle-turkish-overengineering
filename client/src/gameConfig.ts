import { LetterStatus } from './types/game'

export const WORD_LENGTH: number = +import.meta.env.VITE_WORD_LENGTH
export const MAX_ATTEMPTS: number = +import.meta.env.VITE_MAX_ATTEMPTS
export const KEYBOARD_LAYOUT: string[][] = [
  [ 'E', 'R', 'T', 'Y', 'U', 'I', 'O', 'P', 'Ğ', 'Ü' ],
  [ 'A', 'S', 'D', 'F', 'G', 'H', 'J', 'K', 'L', 'Ş', 'İ' ],
  [ 'ENTER', 'Z', 'C', 'V', 'B', 'N', 'M', 'Ö', 'Ç', 'BACKSPACE' ],
]
export const STATUS_COLOR: Record<LetterStatus, string> = {
  [LetterStatus.CORRECT]: 'bg-correct',
  [LetterStatus.PRESENT]: 'bg-present',
  [LetterStatus.ABSENT]: 'bg-absent',
  [LetterStatus.EMPTY]: 'bg-transparent border-2 border-absent',
}
export const VALID_LETTERS_REGEX = /^[ABCÇDEFGĞHIİJKLMNOÖPRSŞTUÜVYZabcçdefgğhıijklmnoöprsştuüvyz]$/
