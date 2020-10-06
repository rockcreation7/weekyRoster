import axios from "axios"

export default axios.create({ 
  baseURL:
    process.env.NODE_ENV === "development"
      ? // Development
        "http://localhost:9000/api"
      : // Production
        "/api",
})
