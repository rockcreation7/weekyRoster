import axios from "axios"
const protocol = window.location.protocol
const host = window.location.hostname
export default axios.create({
  // Development
  baseURL: protocol + "//" + host + ":9000/api",
  // Production
  // baseURL: "/api",
})
