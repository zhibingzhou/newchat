import {
  post,
  get,
  upload
} from '@/utils/request';

//查询用户表情包服务接口
export const ServeFindUserEmoticon = () => {
  return get('/emoticon/user-emoticon');
}

//查询系统表情包服务接口
export const ServeFindSysEmoticon = () => {
  return get('/emoticon/system-emoticon');
}

//设置用户表情包服务接口
export const ServeSetUserEmoticon = (data) => {
  return post('/emoticon/set-user-emoticon', data);
}

//收藏表情包服务接口
export const ServeCollectEmoticon = (data) => {
  return post('/emoticon/collect-emoticon', data);
}

//移除收藏表情包服务接口
export const ServeDelCollectEmoticon = (data) => {
  return post('/emoticon/del-collect-emoticon', data);
}

//上传表情包服务接口
export const ServeUploadEmoticon = (data) => {
  return upload('/emoticon/upload-emoticon', data);
}
