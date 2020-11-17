FROM golang:alpine as builder
LABEL author="wkk5194" 
LABEL authorMail="wangkaikai2014@foxmail.com" 

WORKDIR /app
ADD . /app
# go mod
RUN go env -w GO111MODULE=on && go env -w GOPROXY=https://goproxy.io,direct
# go build
RUN CGO_ENABLED=0 GOOS=linux go build -o anonymousBoard

# 二次镜像
FROM scratch as final

COPY --from=builder /app/anonymousBoard /
COPY --from=builder /app/views /

# 声明运行时容器暴露的端口
EXPOSE 9001

ENTRYPOINT  ["/anonymousBoard"]