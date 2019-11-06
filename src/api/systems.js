import request from "@/utils/request";

const baseURL = "/systems";

export function GetDatabaseDump(loading) {
  return request({
    url: `${baseURL}/dump`,
    method: "get",
    loading
  });
}
