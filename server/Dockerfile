FROM golang:alpine

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.io,direct

WORKDIR /go/src/newchat
COPY . .
RUN go env && go build -o server .

FROM alpine:latest
LABEL MAINTAINER="SliverHorn@sliver_horn@qq.com"

WORKDIR /go/src/newchat
COPY --from=0 /go/src/newchat/server ./
COPY --from=0 /go/src/newchat/config.yaml ./
COPY --from=0 /go/src/newchat/resource ./resource

EXPOSE 8888

ENTRYPOINT ./server
