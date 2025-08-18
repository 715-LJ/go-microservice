#!/bin/bash

# 设置变量
REPO_DIR="/opt/oae/go-microservice"  # Git 仓库路径
IMAGE_NAME="api/v1"    # Docker 镜像名称
CONTAINER_NAME="api"  # Docker 容器名称

# 进入 Git 仓库目录
cd $REPO_DIR || { echo "Repository not found"; exit 1; }

# 拉取最新代码
echo "Pulling latest code..."
git pull
# 构建 Docker 镜像
echo "Building Docker image..."
docker build -t $IMAGE_NAME . || { echo "Docker build failed"; exit 1; }

# 停止并删除旧容器
if [ "$(docker ps -q -f name=$CONTAINER_NAME)" ]; then
    echo "Stopping and removing old container..."
    docker stop $CONTAINER_NAME
    docker rm $CONTAINER_NAME
fi

# 启动新容器
echo "Starting new container..."
docker run -d -p 9090:8080 --name $CONTAINER_NAME $IMAGE_NAME
echo "Deployment completed successfully."
