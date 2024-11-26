FROM registry.cn-hangzhou.aliyuncs.com/golang/golang:1.20 AS builder

WORKDIR /app
COPY . .

# 设置 Go 环境
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct

# 构建应用
RUN go mod download
RUN go build -o x-project cmd/server/main.go

# 第二阶段：运行
FROM registry.cn-hangzhou.aliyuncs.com/ubuntu/ubuntu:22.04

# 安装基础工具和证书
RUN apt-get update && \
    apt-get install -y ca-certificates tzdata && \
    rm -rf /var/lib/apt/lists/*

WORKDIR /app

# 从构建阶段复制二进制文件
COPY --from=builder /app/x-project .
COPY --from=builder /app/config/config.yml /etc/x-project/config.yml

# 设置时区
ENV TZ=Asia/Shanghai

EXPOSE 18000

ENTRYPOINT ["/app/x-project"]
CMD ["-c", "/etc/x-project/config.yml"]
