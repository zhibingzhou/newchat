import {
  post,
  get,
  upload
} from '@/utils/request';

//获取聊天列表服务接口
export const ServeGetTalkList = (data) => {
  return get('/talk/list', data);
}

//聊天列表创建服务接口
export const ServeCreateTalkList = (data) => {
  return post('/talk/create', data);
}

//删除聊天列表服务接口
export const ServeDeleteTalkList = (data) => {
  return post('/talk/delete', data);
}

//对话列表置顶服务接口
export const ServeTopTalkList = (data) => {
  return post('/talk/topping', data);
}

//清除聊天消息未读数服务接口
export const ServeClearTalkUnreadNum = (data) => {
  return post('/talk/update-unread-num', data);
}

//获取聊天记录服务接口
export const ServeTalkRecords = (data) => {
  return get('/talk/records', data);
}

//撤回消息服务接口
export const ServeRevokeRecords = (data) => {
  return post('/talk/revoke-records', data);
}

//删除消息服务接口
export const ServeRemoveRecords = (data) => {
  return post('/talk/remove-records', data);
}

//转发消息服务接口
export const ServeForwardRecords = (data) => {
  return post('/talk/forward-records', data);
}

//获取转发会话记录详情列表服务接口
export const ServeGetForwardRecords = (data) => {
  return get('/talk/get-forward-records', data);
}

//对话列表置顶服务接口
export const ServeSetNotDisturb = (data) => {
  return post('/talk/set-not-disturb', data);
}

//查找用户聊天记录服务接口
export const ServeFindTalkRecords = (data) => {
  return get('/talk/find-chat-records', data);
}

//搜索用户聊天记录服务接口
export const ServeSearchTalkRecords = (data) => {
  return get('/talk/search-chat-records', data);
}

export const ServeGetRecordsContext = (data) => {
  return get('/talk/get-records-context', data);
}

//发送代码块消息服务接口
export const ServeSendTalkCodeBlock = (data) => {
  return post('/talk/send-code-block', data);
}

//发送聊天文件服务接口
export const ServeSendTalkFile = (data) => {
  return post('/talk/send-file', data);
}

//发送聊天图片服务接口
export const ServeSendTalkImage = (data) => {
  return upload('/talk/send-image', data);
}

//发送表情包服务接口
export const ServeSendEmoticon = (data) => {
  return post('/talk/send-emoticon', data);
}
