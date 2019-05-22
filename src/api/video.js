import request from "@/utils/request";

export function getVideos(params) {
  return request({
    url: `/videos`,
    method: "get",
    params
  });
}
