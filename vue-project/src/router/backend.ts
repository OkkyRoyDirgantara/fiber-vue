import axios from 'axios'
import http from './axios'

class BackendService {
  public async getProducts() {
    const response = await http.get('api/v1/product')
    return response.data
  }

  public async login(data: any) {
    const response = await http.post('api/v1/auth/login', data)
    return response.data
  }
}

export default BackendService
