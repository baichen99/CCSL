import request from "@/utils/request";

const baseURL = "/lexical/videos";

export function GetLexicalVideosList(params, loading) {
  return request({
    url: baseURL,
    method: "get",
    params,
    loading
  });
}

export function CreateLexicalVideo(data) {
  return request({
    url: baseURL,
    method: "post",
    data
  });
}

export function GetLexicalVideo(id) {
  return request({
    url: `${baseURL}/${id}`,
    method: "get"
  });
}

export function UpdateLexicalVideo(id, data) {
  return request({
    url: `${baseURL}/${id}`,
    method: "put",
    data
  });
}

export function DeleteLexicalVideo(id) {
  return request({
    url: `${baseURL}/${id}`,
    method: "delete"
  });
}
