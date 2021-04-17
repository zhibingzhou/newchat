#无需任何镜像为基础
#From scratch

#使用了镜像大小体积只有5MB的alpine镜像
FROM scratch
#在docker的根目录下创建相应的使用目录
RUN mkdir -p /go/app
#设置工作路径
WORKDIR /go/app
#把Dockerfile同级目录所有文件拷贝到上面设置的容器的工作目录中
COPY . .
#暴露容器内部端口
EXPOSE 9090
#入口
ENTRYPOINT ["/go/app/newchat"]