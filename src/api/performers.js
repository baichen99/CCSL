import request from "@/utils/request";

const baseURL = "/performers";

export function GetPerformersList(params) {
  return request({
    url: baseURL,
    method: "get",
    params
  });
}

export function CreatePerformer(data) {
  return request({
    url: baseURL,
    method: "post",
    data
  });
}

export function GetPerformer(id) {
  return request({
    url: `${baseURL}/${id}`,
    method: "get"
  });
}

export function UpdatePerformer(id, data) {
  return request({
    url: `${baseURL}/${id}`,
    method: "put",
    data
  });
}

export function DeletePerformer(id) {
  return request({
    url: `${baseURL}/${id}`,
    method: "delete"
  });
}
