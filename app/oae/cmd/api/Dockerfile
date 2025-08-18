# 构建阶段
FROM golang:1.23 AS builder

# 设置工作目录
WORKDIR /app

# 1. 复制项目根目录的 go.mod 和 go.sum
COPY go.mod go.sum ./
RUN go mod download

# 2. 复制整个项目代码（从根目录）
COPY . .

# 传递构建参数
ARG GO_FILE

# 3. 编译指定入口文件
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/oae-api ${GO_FILE}

# 运行阶段
FROM alpine:latest
RUN apk --no-cache add ca-certificates

# 传递构建参数
ARG GO_ETC
# 4. 复制二进制文件和配置文件
COPY --from=builder /app/oae-api .
COPY ${GO_ETC} ./etc

# 5. 暴露端口
EXPOSE 1002

# 6. 启动服务
CMD ["./oae-api"]