import axios from "axios"

export default axios.create({
  // Development
  baseURL: "http://localhost:9000/api",
  // Production
  // baseURL: "/api",
})
