import axios from "axios";
import store from "@/store";
import router from "@/router";
import { Notification, Loading } from "element-ui";
import { getToken } from "@/utils/tools";

// create an axios instance
const service = axios.create({
  baseURL: process.env.VUE_APP_BASE_API,
  timeout: 5000
});

let loading;
let needLoadingRequestCount = 0;

function showFullScreenLoading() {
  if (needLoadingRequestCount === 0) {
    loading = Loading.service({
      lock: true,
      fullscreen: true,
      text: "努力加载中..."
    });
  }
  needLoadingRequestCount++;
}

function hideFullScreenLoading() {
  if (needLoadingRequestCount <= 0) return;
  needLoadingRequestCount--;
  if (needLoadingRequestCount === 0) {
    loading.close();
  }
}

// request interceptor
service.interceptors.request.use(
  config => {
    if (config.loading === true) {
      showFullScreenLoading();
    }
    const token = getToken() || store.getters.token;
    if (token) {
      config.headers["Authorization"] = "Bearer " + token;
    }
    return config;
  },
  error => {
    hideFullScreenLoading();
    Notification({
      title: error,
      type: "error",
      duration: 3 * 1000
    });
    return Promise.reject(error);
  }
);

// response interceptor
service.interceptors.response.use(
  response => {
    hideFullScreenLoading();
    const res = response.data;
    return res;
  },
  error => {
    hideFullScreenLoading();
    const res = error.response.data;
    const statusCode = error.response.status;
    Notification({
      title: res.message,
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
