import axios from 'axios'

export default axios.create({
  baseURL: import.meta.env.VITE_URL_BACKEND,
  headers: {
    Authorization: `Bearer ${localStorage.getItem('token')}`
  }
})
