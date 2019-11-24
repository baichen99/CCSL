import request from "@/utils/request";

const baseURL = "/systems";

export function GetDatabaseDump(loading) {
  return request({
    url: `${baseURL}/dump`,
    method: "get",
    loading
  });
}

export function GetLoginHistory(params) {
  return request({
    url: `${baseURL}/login`,
    method: "get",
    params
  });
}

export function SaveJsError(error) {
  return request({
    url: `${baseURL}/error`,
    method: "post",
    data: error
  });
}

export function GetJsErrorList(params) {
  return request({
    url: `${baseURL}/error`,
    method: "get",
    params
  });
}

export function GetAppInfo(key) {
  return request({
    url: `${baseURL}/info/${key}`,
    method: "get"
  });
}

export function UpdateAppInfo(key, data) {
  return request({
    url: `${baseURL}/info/${key}`,
    method: "put",
    data: { data: data }
  });
}
