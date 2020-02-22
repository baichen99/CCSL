import request from "@/utils/request";

const baseURL = "/carousels"

export function GetCarouselsList(params) {
  return request({
    url: baseURL,
    method: "get",
    params,
  });
}

export function CreateCarousel(data) {
  return request({
    url: baseURL,
    method: "post",
    data
  });
}

export function UpdateCarousel(id, data) {
  return request({
    url: `${baseURL}/${id}`,
    method: "put",
    data
  });
}

export function DeleteCarousel(id) {
  return request({
    url: `${baseURL}/${id}`,
    method: "delete"
  });
}
