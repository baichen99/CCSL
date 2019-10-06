import request from "@/utils/request";

export function GetLexicalWordsList(params) {
  return request({
    url: "/lexical/words",
    method: "get",
    params
  });
}
