import request from "@/utils/request";

export function GetCarouselsList(params, loading) {
  return request({
    url: "/carousels",
    method: "get",
    params,
    loading
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
