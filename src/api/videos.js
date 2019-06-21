import request from "@/utils/request";

export function getUniversalContrastVideos(params) {
  return request({
    url: `/videos`,
    method: "get",
    params
  });
}
