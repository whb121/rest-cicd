# 使用官方Go镜像
FROM golang:1.21-alpine AS builder

WORKDIR /app

# 复制go模块文件
COPY go.mod go.sum ./
RUN go mod download

# 复制源代码
COPY . .

# 构建应用
RUN CGO_ENABLED=0 GOOS=linux go build -o main main.go

# 使用alpine作为运行时
FROM alpine:latest

# 安装CA证书（HTTPS需要）
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# 从构建阶段复制可执行文件
COPY --from=builder /app/main .

# 暴露端口
EXPOSE 8080

# 运行应用
CMD ["./main"]