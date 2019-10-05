import request from "@/utils/request";

const baseURL = "/signs";

export function getSigns(params) {
  return request({
    url: baseURL,
    method: "get",
    params
  });
}

export function updateSign(id, data) {
  return request({
    url: `${baseURL}/${id}`,
    method: "put",
    data
  });
}

export function deleteSign(id) {
  return request({
    url: `${baseURL}/${id}`,
    method: "delete"
  });
}
