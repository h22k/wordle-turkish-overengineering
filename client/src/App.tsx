import Board from './components/sections/Board.js'
import Keyboard from './components/sections/Keyboard.js'
import { KeyboardProvider } from './context/keyboardContext.js'
import { ToastContainer } from 'react-toastify'
import * as React from 'react'

function App() {
  return (
    <KeyboardProvider>
      <div className="flex items-center flex-col gap-3 justify-center p-5">
        <ToastContainer aria-label={ undefined }/>
        <Board/>
        <Keyboard/>
      </div>
    </KeyboardProvider>
  )
}

export default App
