import request from "@/utils/request";

export function Login(data) {
  return request({
    url: "/users/login",
    method: "post",
    data
  });
}

export function RefreshToken() {
  return request({
    url: "/users/refresh",
    method: "get"
  });
}

export function GetUser(userID) {
  return request({
    url: `/users/${userID}`,
    method: "get"
  });
}


