import http from "../http-common"

class RosterDataService {
  getAll() {
    return http.get(`/roster/get`)
  }

  get(date) {
    return http.get(`/roster/get/${date}`)
  }

  create(data) {
    return http.post(`/roster/insert`, data)
  }

  update(date, data) {
    return http.put(`/roster/update/${date}`, data)
  }
}

export default new RosterDataService()
