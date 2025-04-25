export enum LetterStatus {
  EMPTY = 'empty',
  CORRECT = 'correct',
  PRESENT = 'present',
  ABSENT = 'absent',
}

export interface KeyboardContextType {
  letters: Letter[][]
  handleClick: (value: string) => void
  handleChange: (rowIndex: number, index: number, value: string) => void
  currentRow: number
  moveToNextRow: () => void
}

export interface LetterProps {
  status: LetterStatus
  char: string
}

export interface LetterRowProps {
  rowIndex: number
  isFirstRow: boolean
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
}
