import request from "@/utils/request";

const baseURL = "/replies";

export function GetRepliesList(params) {
  return request({
    url: baseURL,
    method: "get",
    params,
  });
}

export function CreateReply(data) {
  return request({
    url: baseURL,
    method: "post",
    data
  });
}

export function GetReply(id) {
  return request({
    url: `${baseURL}/${id}`,
    method: "get"
  });
}

export function UpdateReply(id, data) {
  return request({
    url: `${baseURL}/${id}`,
    method: "put",
    data
  });
}

export function DeleteReply(id) {
  return request({
    url: `${baseURL}/${id}`,
    method: "delete"
  });
}
