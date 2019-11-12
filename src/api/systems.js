import request from "@/utils/request";

const baseURL = "/systems";

export function GetDatabaseDump(loading) {
  return request({
    url: `${baseURL}/dump`,
    method: "get",
    loading
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
