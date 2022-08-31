# deamon server docker file
# docker run -p 10801:10801 --privileged=true -v /var/run/docker.sock:/var/run/docker.sock -v /usr/bin/docker:/usr/bin/docker --name tiandi-deamon -d tiandi-deamon:latest
FROM golang:alpine as builder
#FROM go-gcc:latest as builder

#WORKDIR /go/src/gitee.com/slclub/aliyun-oss
WORKDIR /go/src/tiandi/deamon
COPY ./server/ .
#RUN apk add build-base

RUN go env -w GO111MODULE=on\
	&& go env -w GOPROXY=https://goproxy.cn,direct\
	&& go env -w CGO_ENABLED=1\
	&& go env\
	&& go mod tidy\
	&& go build -o tiandi-deamon .

FROM alpine:latest
LABEL MAINTAINER="slclub@163.com"

WORKDIR /go/src/tiandi/deamon

COPY --from=0 /go/src/tiandi/deamon/tiandi-deamon ./tiandi-deamon
COPY --from=0 /go/src/tiandi/deamon/conf  ./conf
COPY --from=0 /go/src/tiandi/deamon/go.mod ./

RUN mkdir -p ./logs
RUN ls -a

# 同步时间
ENV TZ=Asia/Shanghai
#RUN echo "Asia/shanghai" > /etc/timezone
RUN echo "http://mirrors.aliyun.com/alpine/v3.4/main/" > /etc/apk/repositories \
    && apk --no-cache add tzdata zeromq \
    && ln -snf /usr/share/zoneinfo/$TZ /etc/localtime \
    && echo '$TZ' > /etc/timezone

EXPOSE 10801

ENTRYPOINT ./tiandi-deamon
