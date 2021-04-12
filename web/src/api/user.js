import {
  post,
  get
} from '@/utils/request';

//登录服务接口
export const ServeLogin = (data) => {
  return post('/auth/login', data);
}

// @Summary 获取验证码
export const Captcha = (data) => {
  return post('/auth/captcha', data);
}

// @Summary websocket 注册
export const RegisterWebsocket = (data) => {
  return post('/auth/register_websocket', data);
}

//注册服务接口
export const ServeRegister = (data) => {
  return post('/auth/register', data);
}

//退出登录服务接口
export const ServeLogout = (data) => {
  return post('/auth/logout', data);
}

//刷新登录Token服务接口
export const ServeRefreshToken = (data) => {
  return post('/auth/refresh-token');
}

//修改密码服务接口
export const ServeUpdatePassword = (data) => {
  return post('/users/change-password', data);
}

//修改手机号服务接口
export const ServeUpdateMobile = (data) => {
  return post('/users/change-mobile', data);
}

//修改手机号服务接口
export const ServeUpdateEmail = (data) => {
  return post('/users/change-email', data);
}

//发送手机号修改验证码服务接口
export const ServeSendMobileCode = (data) => {
  return post('/users/send-mobile-code', data);
}

//发送找回密码验证码
export const ServeSendVerifyCode = (data) => {
  return post('/auth/send-verify-code', data);
}

//找回密码服务
export const ServeForgetPassword = (data) => {
  return post('/auth/forget-password', data);
}

//搜索用户信息服务接口
export const ServeSearchUser = (data) => {
  return post('/users/search-user', data);
}

//修改个人信息服务接口
export const ServeUpdateUserDetail = (data) => {
  return post('/users/edit-user-detail', data);
}

//查询用户信息服务接口
export const ServeGetUserDetail = () => {
  return get('/users/detail');
}

//获取用户相关设置信息
export const ServeGetUserSetting = () => {
  return get('/users/setting');
}

//发送邮箱验证码服务接口
export const ServeSendEmailCode = (data) => {
  return post('/users/send-change-email-code', data);
}