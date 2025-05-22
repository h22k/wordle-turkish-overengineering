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
  return newLetters.flat().reduce(
    (acc, { char, status }) => {
      if (!char) return acc

      const upperChar = char.toLocaleUpperCase('tr')
      const prevStatus = acc[upperChar] ?? 'empty'

      if (priority[status] > priority[prevStatus]) {
        acc[upperChar] = status
      }

      return acc
    },
    { ...prevAlphabet }
  )
}
