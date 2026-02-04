# 使用阿里云公共镜像 - 100%可用
FROM registry.cn-hangzhou.aliyuncs.com/google_containers/golang:1.21-alpine AS builder

WORKDIR /app

# 设置Go国内代理
ENV GOPROXY=https://goproxy.cn,direct

# 复制依赖文件
COPY go.mod go.sum ./
RUN go mod download

# 复制源代码
COPY . .

# 构建应用（你的main.go在根目录）
RUN CGO_ENABLED=0 GOOS=linux go build -o app main.go

# 使用阿里云公共的alpine镜像
FROM registry.cn-hangzhou.aliyuncs.com/google_containers/alpine:3.19

# 安装必要包
RUN apk --no-cache add ca-certificates

WORKDIR /app

# 从构建阶段复制
COPY --from=builder /app/app .

# 非root用户运行（安全）
USER 1000

# 暴露端口
EXPOSE 8080

# 启动命令
CMD ["./app"]