import {
  post,
  get,
  upload
} from '@/utils/request';

//上传头像裁剪图片服务接口
export const ServeUploadFileStream = (data) => {
  return post('/upload/file-stream', data);
}

//查询大文件拆分信息服务接口
export const ServeFindFileSplitInfo = (data = {}) => {
  return get('/upload/get-file-split-info', data);
}

//文件拆分上传服务接口
export const ServeFileSubareaUpload = (data = {}, options = {}) => {
  return upload('/upload/file-subarea-upload', data, options);
}
