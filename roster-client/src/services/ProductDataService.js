import http from "../http-common"

class ProductDataService {
  getAll() {
    return http.get(`/product/get`)
  }

  get(date) {
    return http.get(`/product/get/${date}`)
  }

  create(data) {
    return http.post(`/product/insert`, data)
  }

  update(date, data) {
    return http.put(`/product/update/${date}`, data)
  }
}

// why new() here
export default new ProductDataService()
