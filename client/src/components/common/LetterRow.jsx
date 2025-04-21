import LetterBox from './LetterBox.jsx'

function LetterRow({ letters }) {
  const rowLetters = Array.from({ length: 5 }).map((_, index) => {
    return letters[index] || { char: "", status: "empty" };
  });

  return (
    <div className="flex gap-1.5">
      {rowLetters.map((letter, index) => (
        <LetterBox key={index} letter={letter.char} status={letter.status} />
      ))}
    </div>
  );
}

export default LetterRow;
