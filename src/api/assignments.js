import request from "@/utils/request";

const baseURL = "/assignments"

export function GetAssignmentsList(params) {
    return request({
        url: baseURL,
        method: "get",
        params
    })
}

export function CreateAssignment(data) {
    return request({
        url: baseURL,
        method: "post",
        data
    })
}

export function GetAssignment(id) {
    return request({
        url: `${baseURL}/${id}`,
        method: "get"
    })
}
export function UpdateAssignment(id, data) {
    return request({
        url: `${baseURL}/${id}`,
        method: "put",
        data
    })
}

export function DeleteAssignment(id) {
    return request({
        url: `${baseURL}/${id}`,
        method: "delete"
    })
}