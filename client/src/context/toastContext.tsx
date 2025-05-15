import React, { createContext, useContext } from 'react'
import { cssTransition, toast, ToastContainer } from 'react-toastify'

const ToastContext = createContext({
  notify: (msg: string) => {
  },
})

export const ToastProvider: React.FC<{ children: React.ReactNode }> = ({ children }) => {
  const notify = (msg: string) => {
    toast.error(msg, {
      position: 'top-center',
      autoClose: 1000,
      hideProgressBar: true,
      closeOnClick: false,
      closeButton: false,
      pauseOnHover: true,
      className: 'text-sm text-black font-semibold rounded-md shadow-md px-2 py-1 min-w-0 w-auto! min-h-0!',
      icon: false,
    })
  }

  return (
    <ToastContext.Provider value={ { notify } }>
      <ToastContainer
        transition={ cssTransition({
          enter: 'none',
          exit: 'fadeOut',
        }) }
      />
      { children }
    </ToastContext.Provider>
  )
}

export const useToast = () => useContext(ToastContext)
