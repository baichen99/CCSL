import request from "@/utils/request";

const baseURL = "/files";

export function UploadFile(formData, dir) {
    return request({
        url: baseURL,
        method: "post",
        params: { dir },
        data: formData,
        headers: {
            "Content-Type": "multipart/form-data"
        }
    });
}