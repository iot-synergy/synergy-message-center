FROM alpine:3.19

# Define the project name | 定义项目名称
ARG PROJECT=mcms
# Define the config file name | 定义配置文件名
ARG CONFIG_FILE=mcms.yaml

WORKDIR /app
ENV PROJECT=${PROJECT}
ENV CONFIG_FILE=${CONFIG_FILE}

ENV TZ=Asia/Shanghai
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk update --no-cache && apk add --no-cache tzdata

COPY /app/mcms-rpc ./
COPY /app/etc/${CONFIG_FILE} ./etc/

EXPOSE 9106

ENTRYPOINT ./mcms-rpc -f etc/${CONFIG_FILE}