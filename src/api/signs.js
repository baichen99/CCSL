import request from "@/utils/request";

const baseURL = "/signs";

export function GetSignsList(params) {
  return request({
    url: baseURL,
    method: "get",
    params
  });
}

export function CreateSign(data) {
  return request({
    url: baseURL,
    method: "post",
    data
  });
}

export function GetSign(id) {
  return request({
    url: `${baseURL}/${id}`,
    method: "get"
  });
}

export function UpdateSign(id, data) {
  return request({
    url: `${baseURL}/${id}`,
    method: "put",
    data
  });
}

export function DeleteSign(id) {
  return request({
    url: `${baseURL}/${id}`,
    method: "delete"
  });
}
