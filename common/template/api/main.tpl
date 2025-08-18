package main

import (
	"flag"
	"fmt"

	{{.importPackages}}
)

var configFile = flag.String("f", "etc/{{.serviceName}}.yaml", "the config file")

func main() {
	flag.Parse()

    // 加载配置
    var c config.Config
    conf.MustLoad(*configFile, &c, conf.UseEnv())

    basic.Init()

    // 服务初始化
    server := rest.MustNewServer(c.RestConf, rest.WithUnauthorizedCallback(result.JwtUnauthorizedCallback), rest.WithUnsignedCallback(result.UnsignedCallback), rest.WithCors())
    defer server.Stop()

    // 添加 ErrorHandlingMiddleware
    server.Use(middleware.ErrorHandlingMiddleware)

    ctx := svc.NewServiceContext(c)
    handler.RegisterHandlers(server, ctx)

    // 添加控制台输出
    logx.AddWriter(logx.NewWriter(os.Stdout))

    fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
    server.Start()
}
