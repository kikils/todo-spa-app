import axios from "axios";
import router from "./router";

const baseURL = process.env.VUE_APP_API_BASE_URL;

const api = axios.create({
  baseURL,
  headers: {
    Authorization: `Bearer ${localStorage.getItem("jwt")}`,
  },
});

api.interceptors.response.use(
  (response) => {
    return response;
  },
  (error) => {
    if (error.response.status === 401) {
      router.push("/signin").catch((err) => console.log(err));
    }
    return error;
  }
);

export default api;
