# Lumen IM 即时聊天系统(前端)

### 1. 项目介绍
Lumen IM 是一个网页版在线即时聊天项目，前端使用 Element-ui + Vue，后端采用了基于 Swoole 开发的 Hyperf 协程框架进行接口开发，并使用 WebSocket 服务进行消息实时推送。目前后端 WebSocket 已支持分布式集群部署。

目前该项目是在 [旧版本](https://github.com/gzydong/LumenIM/tree/v1.0.0) 项目的基础上进行了后端重构，且前后端都有较大的改动。

### 2. 功能模块
- 基于 Swoole WebSocket 服务做消息即时推送
- 支持私聊及群聊
- 支持多种聊天消息类型 例如:文本、代码块、图片及其它类型文件，并支持文件下载
- 支持聊天消息撤回、删除(批量删除)、转发消息(逐条转发、合并转发)
- 支持编写个人笔记、支持笔记分享(好友或群)

### 3. 项目预览
- 地址： [http://im.gzydong.club](http://im.gzydong.club)
- 账号： 18798272054 或 18798272055
- 密码： admin123

### 4. 项目安装及部署
```bash
## 克隆项目源码包
git clone git@github.com:gzydong/LumenIM.git

## 安装项目依赖扩展组件
npm install

# 启动本地开发环境
npm run serve

## 生产环境构建项目
npm run build

## 生产环境构建项目并查看构建报告
npm run build --report
```

#### 修改 .env 配置信息

```env
VUE_APP_API_BASE_URL=http://xxx.yourdomain.com
VUE_APP_WEB_SOCKET_URL=ws://xxx.yourdomain.com/socket.io
VUE_APP_WEBSITE_NAME="Lumen IM"
```

#### 关于 Nginx 的一些配置
```nginx
server {
    listen       80;
    server_name  www.yourdomain.com;

    root /project-path/dist;
    index  index.html;

    ## 解决 VueRouter History 模式下 页面刷新404问题
    location / {
      try_files $uri $uri/ /index.html;
    }

    location ~ .*\.(gif|jpg|jpeg|png|bmp|swf|flv|ico)$ {
        expires 30d;
    }

    location ~ .*\.(js|css)?$ {
        expires 7d;
    }
}
```
@link 后端源码 [https://github.com/gzydong/hyperf-chat](https://github.com/gzydong/hyperf-chat)

注意：项目需要与后端一起使用，目前后端源码还未开源，如有需要可联系 837215079@qq.com

#### 如果你觉得还不错，请 Star , Fork 给作者鼓励一下。

## License

[LICENSE](LICENSE)

