import request from "@/utils/request";

const baseURL = "/news";

export function GetNewsList() {
  return request({
    url: baseURL,
    method: "get"
  });
}

export function CreateNews(data) {
  return request({
    url: baseURL,
    method: "post",
    data
  });
}

export function UpdateNews(id, data) {
  return request({
    url: `${baseURL}/${id}`,
    method: "put",
    data
  });
}

export function DeleteNews(id) {
  return request({
    url: `${baseURL}/${id}`,
    method: "delete"
  });
}
