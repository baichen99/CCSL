import request from "@/utils/request";

export function GetCarouselList(params) {
  return request({
    url: "/carousels",
    method: "get",
    params
  });
}

export function CreateCarousel(data) {
  return request({
    url: "/carousels",
    method: "post",
    data
  });
}

export function UpdateCarousel(id, data) {
  return request({
    url: `/carousels/${id}`,
    method: "put",
    data
  });
}

export function DeleteCarousel(id) {
  return request({
    url: `/carousels/${id}`,
    method: "delete"
  });
}
