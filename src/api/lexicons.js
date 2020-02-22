import request from "@/utils/request";

const baseURL = "/lexicons";

export function GetWordsList(params) {
  return request({
    url: baseURL,
    method: "get",
    params,
  });
}

export function CreateWord(data) {
  return request({
    url: baseURL,
    method: "post",
    data
  });
}

export function GetWord(id) {
  return request({
    url: `${baseURL}/${id}`,
    method: "get"
  });
}

export function UpdateWord(id, data) {
  return request({
    url: `${baseURL}/${id}`,
    method: "put",
    data
  });
}

export function DeleteWord(id) {
  return request({
    url: `${baseURL}/${id}`,
    method: "delete"
  });
}
