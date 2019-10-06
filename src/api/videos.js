import request from "@/utils/request";

export function GetLexicalVideosList(params) {
  return request({
    url: "/lexical/videos",
    method: "get",
    params
    //loading:true
  });
}
