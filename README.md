### go-microservice

- go-zero
- nginx网关
- elasticsearch
- docker
- docker-compose
- mysql
- redis

项目目录结构如下：

- app：所有业务代码包含api、rpc以及mq（消息队列、延迟队列、定时任务）
    - mesas：微服务
        - cmd：可执行层代码
            - api： HTTP服务
                - desc：api定义文件
                - etc：系统配置文件
                - internal：业务层代码
                    - config：配置中心
                    - handler：请求处理
                    - logic：逻辑层代码
                    - svc：上下文配置
                    - types：系统定义结构体
                - mesas.go：
                - Dockerfile：镜像构建脚本
        - model：模型层代码
    - oae
- common：
    - basic：系统初始化配置，可以自定义初始化不同微服务的链接池
        - mysql
        - redis
    - error：错误处理
    - middleware：中间件
    - result：自定义接口返回值结构信息
    - template：配置脚手架模板
    - translator：错误翻译
    - xerr：错误码及错误信息配置
- data：临时文件
- deploy：
    - nginx：网关配置
    - script：
        - mysql：生成model的sh工具
- docker-compose.yml：容器脚本
- go.mod：模块依赖
- go.sum：模块依赖
- Makefile：执行命令脚本
- README.md：描述

## 网关

nginx做对外网关

## 开发模式

本项目使用的是微服务开发，api （http） + rpc（grpc） ， api充当聚合服务，复杂、涉及到其他业务调用的统一写在rpc中，如果一些不会被其他服务依赖使用的简单业务，可以直接写在api的logic中

## 部署

本项目开发环境推荐docker-compose

gitlab + jenkins + harbor + k8s

在jenkins中点击部署对应的服务，会去gitlab拉取代码-->再去拉取线上配置（线上配置单独一个git库，为什么不用配置中心，部署文档中有介绍）---->自动构建镜像-->推送到harbor镜像仓库--->
使用kubectl自动发布到k8s中---->前面要挂一个nginx做网关统一入口

### 项目环境搭建

#### 1、clone代码&更新依赖

```shell
$ git clone https://github.com/715-LJ/go-microservice.git
$ go mod tidy
```

本项目采用modd热加载功即时修改代码及时生效，并且不需要每次都要重启，改了代码自动就在容器中重新加载了，本地不需要启动服务。

```shell
go install github.com/cortesi/modd/cmd/modd@latest
modd.conf
app/mesas/cmd/api/**/*.go {
    prep: go build -o data/server/mesas-api  -v app/mesas/cmd/api/mesas.go
    daemon +sigkill: ./data/server/mesas-api -f app/mesas/cmd/api/etc/mesas-api.yaml
}
modd
```

#### 2、启动服务

##### 2.1 拉取运行环境镜像

因为本项目是用docker+热加载，即改即生效

前台app下所有api+rpc服务统一使用modd + golang

直接docker-compose去启动可以，但是考虑依赖可能会比较大，会影响启动项目，所以最好先把这个镜像拉取下来再去启动项目

##### 2.2 启动项目

```shell
$ docker-compose up -d 
```

【注】依赖的是项目根目录下的docker-compose.yml配置

