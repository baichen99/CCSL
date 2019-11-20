import request from "@/utils/request";

const baseURL = "/members";

export function GetMembersList(params, loading) {
  return request({
    url: baseURL,
    method: "get",
    params,
    loading
  });
}

export function GetMember(id) {
  return request({
    url: `${baseURL}/${id}`,
    method: "get"
  });
}

export function CreateMember(data) {
  return request({
    url: baseURL,
    method: "post",
    data
  });
}

export function UpdateMember(id, data) {
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
