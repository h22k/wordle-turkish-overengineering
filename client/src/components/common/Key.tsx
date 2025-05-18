import { IoBackspaceOutline } from 'react-icons/io5'
import { KeyProps } from '../../types/game'

function Key({ value, onClick }: KeyProps) {
  return (
    <button
      onClick={() => onClick(value)}
      className={`rounded-md px-3 py-2 ${value === 'ENTER' ? 'text-xs' : 'text-xl'} focus:outline-none font-bold text-white bg-keyboard min-w-[43px] h-[58px] flex items-center justify-center`}
    >
      {value === 'BACKSPACE' ? <IoBackspaceOutline size={26} /> : value}
    </button>
  )
}

export default Key
