import request from "@/utils/request";

const baseURL = "/posts";

export function GetPostsList(params) {
  return request({
    url: baseURL,
    method: "get",
    params,
  });
}

export function CreatePost(data) {
  return request({
    url: baseURL,
    method: "post",
    data
  });
}

export function GetPost(id) {
  return request({
    url: `${baseURL}/${id}`,
    method: "get"
  });
}

export function UpdatePost(id, data) {
  return request({
    url: `${baseURL}/${id}`,
    method: "put",
    data
  });
}

export function DeletePost(id) {
  return request({
    url: `${baseURL}/${id}`,
    method: "delete"
  });
}
