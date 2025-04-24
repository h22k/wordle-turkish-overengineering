import Board from './components/sections/Board.jsx'
import Keyboard from './components/sections/Keyboard.jsx'
import { KeyboardProvider } from './context/keyboardContext.jsx'
import { ToastContainer } from 'react-toastify'

function App() {
  return (
    <KeyboardProvider>
      <div className="flex items-center flex-col gap-3 justify-center p-5">
        <ToastContainer/>
        <Board/>
        <Keyboard/>
      </div>
    </KeyboardProvider>
  )
}

export default App
