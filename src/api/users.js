import request from "@/utils/request";

const baseURL = "/users";

export function Login(data) {
  return request({
    url: `${baseURL}/login`,
    method: "post",
    data
  });
}

export function RefreshToken() {
  return request({
    url: `${baseURL}/refresh`,
    method: "get"
  });
}

export function GetUser(id) {
  return request({
    url: `${baseURL}/${id}`,
    method: "get"
  });
}

export function GetUsersList(params, loading) {
  return request({
    url: baseURL,
    method: "get",
    params,
    loading
  });
}

export function CreateUser(data) {
  return request({
    url: baseURL,
    method: "post",
    data
  });
}

export function UpdateUser(id, data) {
  return request({
    url: `${baseURL}/${id}`,
    method: "put",
    data
  });
}

export function DeleteUser(id) {
  return request({
    url: `${baseURL}/${id}`,
    method: "delete"
  });
}
