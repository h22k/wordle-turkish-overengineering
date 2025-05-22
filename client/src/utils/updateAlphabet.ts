import { LetterProps, LetterStatus } from '../types/game'

const priority: Record<LetterStatus, number> = {
  correct: 2,
  present: 1,
  absent: 0,
  empty: -1,
}

export const getUpdatedAlphabet = (
  prevAlphabet: Record<string, LetterStatus>,
  newLetters: LetterProps[][]
): Record<string, LetterStatus> => {
  const updated = { ...prevAlphabet }

  for (const row of newLetters) {
    for (const { char, status } of row) {
      if (!char) continue
      const upperChar = char.toLocaleUpperCase('tr')

      if (!(upperChar in updated) || priority[status] > priority[updated[upperChar]]) {
        updated[upperChar] = status
      }
    }
  }

  return updated
}
