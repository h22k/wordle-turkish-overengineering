function LetterBox({ letter, status }) {
  const statusColor = {
    correct: 'bg-[#538d4e] text-white',
    present: 'bg-[#b59f3b] text-white',
    absent: 'bg-[#3a3a3c] text-white',
    empty: 'bg-transparent border-2 border-[#3a3a3c] text-white',
  }

  return (
    <div
      className={ `w-[52px] h-[52px] flex items-center justify-center text-xl font-bold uppercase ${ statusColor[status] }` }
    >
      { letter }
    </div>
  )
}

export default LetterBox
