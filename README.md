### go-microservice

- k8s
- go-zero
- nginx网关
- elasticsearch
- go-queue
- docker
- docker-compose
- mysql
- redis
- jenkins
- gitlab

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
- deploy：
    - nginx：网关配置
    - script：
        - mysql：生成model的sh工具
- deploy.sh 更新脚本
- docker-compose.yml：容器脚本
- Dockerfile：镜像构建脚本
- go.mod：模块依赖
- go.sum：模块依赖
- Makefile：执行命令脚本
- README.md：描述
- update.sh：更新脚本

## 网关

nginx做对外网关

## 开发模式

本项目使用的是微服务开发，api （http） + rpc（grpc） ， api充当聚合服务，复杂、涉及到其他业务调用的统一写在rpc中，如果一些不会被其他服务依赖使用的简单业务，可以直接写在api的logic中

## 日志

关于日志，统一使用filebeat收集，上报到kafka中，由于logstash懂得都懂，资源占用太夸张了，这里使用了go-stash替换了logstash

链接：https://github.com/kevwan/go-stash
go-stash是由go-zero开发团队开发的，性能很高不占资源，主要代码量没多少，只需要配置就可以使用，很简单。它是把kafka数据源同步到elasticsearch中，默认不支持elasticsearch账号密码，我fork了一份修改了一下，很简单支持了账号、密码

## 发布订阅

发布订阅使用的是go-zero开发团队开发的go-queue， 链接：https://github.com/zeromicro/go-queue

## 消息队列、延迟队列、定时任务

消息队列、延迟队列、定时任务本项目使用的是asynq ，基于redis开发的简单中间件，

当然，消息队列你也可以使用go-queue

链接：https://github.com/hibiken/asynq

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

本项目采用modd热加载功即时修改代码及时生效，并且不需要每次都要重启，改了代码自动就在容器中重新加载了，本地不需要启动服务，本地安装的sdk就是写代码自动提示使用的，实际运行是以来容器中 lyumikael/go-modd-env:
v1.0.0的golang环境。所以使用goland、vscode都一样

#### 2、启动服务

##### 2.1 拉取运行环境镜像

因为本项目是用docker+热加载，即改即生效

前台app下所有api+rpc服务统一使用modd + golang

直接docker-compose去启动可以，但是考虑依赖可能会比较大，会影响启动项目，所以最好先把这个镜像拉取下来再去启动项目

```shell
$ docker pull lyumikael/gomodd:v1.20.3 #这个是app下所有的api+rpc启动服务使用的，如果你是 "mac m1" : lyumikael/go-modd-env:v1.0.0
```

【注】后续如果app下新增业务，要记得在项目根目录下的modd.conf复制添加一份就可以了

​ 关于modd更多用法可以去这里了解 ： https://github.com/cortesi/modd ， 本项目镜像只是将golang-1.17.7-alpine作为基础镜像安装了modd在内部，

如果你想把goctl、protoc、golint等加进去，不用我的镜像直接制作一个镜像也一样的哈

##### 2.2 启动项目

```shell
$ docker-compose up -d 
```

【注】依赖的是项目根目录下的docker-compose.yml配置

