import http from "../http-common"

class ProductDataService {
  getAll() {
    return http.get(`/product/get`)
  }

  get(id) {
    return http.get(`/product/get/${id}`)
  }

  create(id) {
    return http.post(`/product/insert`, id)
  }

  update(id, data) {
    return http.put(`/product/update/${id}`, data)
  }

  delete(id) {
    return http.delete(`/product/delete/${id}`)
  }
}

// why new() here
export default new ProductDataService()
