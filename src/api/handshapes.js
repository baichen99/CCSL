import request from "@/utils/request";

const baseURL = "/handshapes";

export function GetHandshapesList(params) {
  return request({
    url: baseURL,
    method: "get",
    params,
  });
}

export function CreateHandshape(data) {
  return request({
    url: baseURL,
    method: "post",
    data
  });
}

export function GetHandshape(id) {
  return request({
    url: `${baseURL}/${id}`,
    method: "get"
  });
}

export function UpdateHandshape(id, data) {
  return request({
    url: `${baseURL}/${id}`,
    method: "put",
    data
  });
}

export function DeleteHandshape(id) {
  return request({
    url: `${baseURL}/${id}`,
    method: "delete"
  });
}
