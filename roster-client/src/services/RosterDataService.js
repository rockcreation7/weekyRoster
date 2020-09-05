import http from "../http-common"

class RosterDataService {
  getAll() {
    return http.get("/get")
  }

  get(date) {
    return http.get(`/get/${date}`)
  }

  create(data) {
    return http.post("/insert", data)
  }

  update(date, data) {
    return http.put(`/update/${date}`, data)
  }
  /* 
  delete(id) {
    return http.delete(`/tutorials/${id}`);
  }

  deleteAll() {
    return http.delete(`/tutorials`);
  }

  findByTitle(title) {
    return http.get(`/tutorials?title=${title}`);
  } */
}

export default new RosterDataService()
