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
  moveToNextRow: () => void
  activeBoxIndex: number
  setActiveBoxIndex: (index: number) => void
  shakeRowIndex: number | null
  triggerShake: (row: number) => void
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
