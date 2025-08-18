### 一、项目简介

项目地址 :  http://172.16.40.201:1080/ucenter/go-microservice

原型图地址 : https://mastergo.com/file/107093281047006?fileOpenFrom=project&page_id=17%3A682&shareId=107093281047006

整个项目使用了go-zero开发的微服务，基本包含了go-zero以及gorm等一些中间件

项目目录结构如下：

- app：所有业务代码包含HTTP等中间件
    - api：API层
- common：通用组件 basic、error、middleware、result、template、translator、xerr 等
- deploy：
    - script：
        - gencode：生成api、rpc，以及创建kafka语句，复制粘贴使用
        - mysql：生成model的sh工具
    - goctl: 该项目goctl的template，goctl生成自定义代码模版，template用法可参考go-zero文档，复制到家目录下.goctl即可

### 二、用到技术栈

- go-zero

- mysql

- redis
