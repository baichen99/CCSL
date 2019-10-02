import request from "@/utils/request";

export function getLexicalWords(params) {
  return request({
    url: "/lexical/words",
    method: "get",
    params
  });
}
