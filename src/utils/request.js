import axios from "axios";
import store from "@/store";
import router from "@/router";
import { Message } from "element-ui";
import { getToken } from "@/utils/tools";

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
      // If returns 401, then clean user token and force login
      case 401:
        store.dispatch("user/logout");
        router.push("/login");
        break;

      default:
        break;
    }
    return Promise.reject(error);
  }
);

export default service;
