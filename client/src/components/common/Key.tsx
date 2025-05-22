import { IoBackspaceOutline } from 'react-icons/io5'
import { KeyProps } from '../../types/game'
import { useKeyboard } from '../../hooks/useKeyboard'
import { STATUS_COLOR } from '../../gameConfig'

function Key({ value, onClick }: KeyProps) {
  const { alphabet } = useKeyboard()

  return (
    <button
      onClick={() => onClick(value)}
      className={`rounded-md sm:px-3 sm:py-2 
      ${value === 'ENTER' ? 'text-xs' : 'text-sm sm:text-xl'} 
      focus:outline-none font-bold text-white sm:min-w-[43px] h-[58px] flex flex-auto items-center justify-center
      ${
        alphabet[value.toLocaleUpperCase('tr')]
          ? STATUS_COLOR[alphabet[value.toLocaleUpperCase('tr')]]
          : 'bg-keyboard'
      } 
        `}
    >
      {value === 'BACKSPACE' ? <IoBackspaceOutline size={26} /> : value}
    </button>
  )
}

export default Key
