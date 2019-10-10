import request from "@/utils/request";

const baseURL = "/lexical/words";

export function GetLexicalWordsList(params) {
  return request({
    url: baseURL,
    method: "get",
    params
  });
}

export function CreateLexicalWord(data) {
  return request({
    url: baseURL,
    method: "post",
    data
  });
}

export function GetLexicalWord(id) {
  return request({
    url: `${baseURL}/${id}`,
    method: "get"
  });
}

export function UpdateLexicalWord(id, data) {
  return request({
    url: `${baseURL}/${id}`,
    method: "put",
    data
  });
}

export function DeleteLexicalWord(id) {
  return request({
    url: `${baseURL}/${id}`,
    method: "delete"
  });
}
