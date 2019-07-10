import request from "@/utils/request";

export function GetNewsList() {
  return request({
    url: "/news",
    method: "get"
  });
}

export function CreateNews(data) {
  return request({
    url: "/news",
    method: "post",
    data
  });
}

export function UpdateNews(id, data) {
  return request({
    url: `/news/${id}`,
    method: "put",
    data
  });
}

export function DeleteNews(id) {
  return request({
    url: `/news/${id}`,
    method: "delete"
  });
}
