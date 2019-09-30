import axios from "axios";
import store from "@/store";
import router from "@/router";
import { Message } from "element-ui";
import { getToken, removeToken, removeUser } from "@/utils/tools";

// create an axios instance
const service = axios.create({
  baseURL: process.env.VUE_APP_BASE_API,
  timeout: 5000
});

// request interceptor
service.interceptors.request.use(
  config => {
    const token = getToken() || store.getters.token;
    if (token) {
      config.headers["Authorization"] = "Bearer " + token;
    }
    return config;
  },
  error => {
    Message({
      message: error,
      type: "error",
      duration: 5 * 1000
    });
    return Promise.reject(error);
  }
);

// response interceptor
service.interceptors.response.use(
  response => {
    const res = response.data;
    return res;
  },
  error => {
    const res = error.response.data;
    const statusCode = error.response.status;
    Message({
      message: res.error,
      type: "error",
      duration: 5 * 1000
    });
    switch (statusCode) {
      case 401:
        removeToken();
        removeUser();
        router.push("/login");
        break;

      default:
        break;
    }
    return Promise.reject(error);
  }
);

export default service;
