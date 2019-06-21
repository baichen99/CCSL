import request from "@/utils/request";

export function getWords(params) {
  return request({
    url: `/words`,
    method: "get",
    params
  });
}
