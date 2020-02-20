import request from "@/utils/request";

const baseURL = "/classes";

export function GetClassesList(params) {
    return request({
        url: baseURL,
        method: "get",
        params
    })
}

export function CreateClass(data) {
    return request({
        url: baseURL,
        method: "post",
        data
    })
}

export function GetClass(id) {
    return request({
        url: `${baseURL}/${id}`,
        method: "get"
    })
}

export function UpdateClass(id, data) {
    return request({
        url: `${baseURL}/${id}`,
        method: "put",
        data
    })
}

export function DeleteClass(id) {
    return request({
        url: `${baseURL}/${id}`,
        method: "delete"
    })
}