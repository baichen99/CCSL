import request from "@/utils/request";

export function getLexicalVideos(params) {
  return request({
    url: "/lexical/videos",
    method: "get",
    params
  });
}
