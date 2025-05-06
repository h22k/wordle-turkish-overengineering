import axios from 'axios'

const axiosClient = axios.create({
  baseURL: import.meta.env.VITE_API_URL_LOCAL,
  headers: {
    'Content-Type': 'application/json',
  },
})

export default axiosClient
