# 使用官方Go镜像作为构建环境
FROM golang:1.21-alpine AS builder

# 设置工作目录
WORKDIR /app

# 复制go.mod和go.sum文件
COPY go.mod go.sum ./

# 下载依赖
RUN go mod download

# 复制源代码
COPY . .

# 构建应用（注意：你的main.go在根目录）
RUN CGO_ENABLED=0 GOOS=linux go build -o main main.go

# 使用更小的运行时镜像
FROM alpine:latest

# 安装ca-certificates（用于HTTPS）
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# 从构建阶段复制可执行文件
COPY --from=builder /app/main .

# 暴露端口（你的应用应该是8080）
EXPOSE 8080

# 运行应用
CMD ["./main"]