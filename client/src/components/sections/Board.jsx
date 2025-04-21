import LetterRow from '../common/LetterRow.jsx'

function Board() {
  const rows = []

  for ( let i = 0; i < 6; i++ ) {
    rows.push(
      <LetterRow letters={ '' }/>,
    )
  }

  return <div className="grid gap-2">{ rows }</div>
}

export default Board
