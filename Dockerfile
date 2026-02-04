# 使用ACR代理的Docker Hub官方镜像（不需要网络访问国外）
FROM crpi-nnlgwphnth3rx1pc.cn-hangzhou.personal.cr.aliyuncs.com/library/golang:1.21-alpine AS builder

WORKDIR /app

# 设置Go代理（国内加速）
ENV GOPROXY=https://goproxy.cn,direct

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main main.go

# 使用ACR代理的alpine镜像
FROM crpi-nnlgwphnth3rx1pc.cn-hangzhou.personal.cr.aliyuncs.com/library/alpine:latest

RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["./main"]