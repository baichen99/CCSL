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

export function GetUser(userID) {
  return request({
    url: `${baseURL}/${userID}`,
    method: "get"
  });
}

export function GetUsersList(params) {
  return request({
    url: baseURL,
    method: "get",
    params
  });
}
