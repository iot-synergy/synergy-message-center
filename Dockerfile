FROM golang:1.21-alpine  as builder

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn
COPY . /app
WORKDIR /app
RUN go mod tidy --compat=1.21 && go build -o /app/mcms-rpc mcms.go

FROM alpine:3.19

# Define the project name | 定义项目名称
ARG PROJECT=mcms
# Define the config file name | 定义配置文件名
ARG CONFIG_FILE=mcms.yaml
# Define the author | 定义作者
ARG AUTHOR="yuansu.china.work@gmail.com"

LABEL org.opencontainers.image.authors=${AUTHOR}

WORKDIR /app
ENV PROJECT=${PROJECT}
ENV CONFIG_FILE=${CONFIG_FILE}

ENV TZ=Asia/Shanghai
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk update --no-cache && apk add --no-cache tzdata

COPY --from=builder /app/mcms-rpc ./
COPY --from=builder /app/etc/${CONFIG_FILE} ./etc/

EXPOSE 9106

ENTRYPOINT ./mcms-rpc -f etc/${CONFIG_FILE}