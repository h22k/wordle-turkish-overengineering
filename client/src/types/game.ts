export enum LetterStatus {
  EMPTY = 'empty',
  CORRECT = 'correct',
  PRESENT = 'present',
  ABSENT = 'absent',
}

export interface LetterProps {
  status: LetterStatus
  char: string
}

export interface KeyboardContextType {
  letters: LetterProps[][]
  handleChange: (value: string) => void
  currentRow: number
  activeBoxIndex: number
  setActiveBoxIndex: (index: number) => void
  shakeRowIndex: number | null
  triggerShake: (row: number) => void
  getFirstEmptyBoxIndex: (row: number) => number
  submittedRow: number | null
  maxAttempts: number
  wordLength: number
  loading: boolean
  alphabet: Record<string, LetterStatus>
}

export interface LetterRowProps {
  rowIndex: number
  currentRow: number
}

export interface KeyProps {
  value: string
  onClick: (value: string) => void
}

export interface LetterBoxProps {
  letter: string
  status: LetterStatus
  onChange?: (value: string) => void
  isFirstBox: boolean
  index?: number
  rowIndex?: number
}

export interface APILetter {
  char: string
  status: LetterStatus
}

export interface APIGuess {
  word: string
  letters: APILetter[]
}
