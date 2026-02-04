# 使用阿里云镜像源的Go镜像（替代docker.io）
FROM registry.cn-hangzhou.aliyuncs.com/aliyunfc/golang:1.21-alpine AS builder

# 设置工作目录
WORKDIR /app

# 配置Go国内代理（加速依赖下载）
RUN go env -w GOPROXY=https://goproxy.cn,direct

# 复制go.mod和go.sum文件
COPY go.mod go.sum ./

# 下载依赖
RUN go mod download

# 复制源代码
COPY . .

# 构建应用（静态编译，无系统依赖）
RUN CGO_ENABLED=0 GOOS=linux go build -o main main.go

# 使用阿里云镜像源的alpine镜像（替代docker.io）
FROM registry.cn-hangzhou.aliyuncs.com/aliyunfc/alpine:3.19

# 安装ca-certificates（用于HTTPS）
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# 从构建阶段复制可执行文件
COPY --from=builder /app/main .

# 赋予执行权限（避免权限问题）
RUN chmod +x ./main

# 暴露端口
EXPOSE 8080

# 运行应用
CMD ["./main"]