#!/bin/bash

# 拉取最新代码
echo "Pulling latest code..."
git pull

# 构建 Docker 镜像
echo "Building Docker image..."
docker-compose build

# 启动新容器
echo "Starting new container..."
docker-compose up -d
echo "Deployment completed successfully."