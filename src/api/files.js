import request from "@/utils/request";

const baseURL = "/files";

export function UploadFile(formData, dir, progress) {
  return request({
    url: baseURL,
    method: "post",
    params: { dir },
    data: formData,
    headers: {
      "Content-Type": "multipart/form-data"
    },
    onUploadProgress: progressEvent => {
      const percentage = progressEvent.loaded / progressEvent.total | 0;
      progress(percentage);
    }
  });
}
