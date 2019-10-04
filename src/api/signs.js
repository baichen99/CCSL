import request from "@/utils/request";

const baseURL = "/signs";

export function getSigns(params) {
  return request({
    url: baseURL,
    method: "get",
    params
  });
}