import request from "@/utils/request";

const baseURL = "/notifications";

export function GetNotificationsList(params) {
  return request({
    url: baseURL,
    method: "get",
    params
  });
}

export function GetNotification(id) {
  return request({
    url: `${baseURL}/${id}`,
    method: "get"
  });
}

export function DeleteNotification(id) {
  return request({
    url: `${baseURL}/${id}`,
    method: "delete"
  });
}
