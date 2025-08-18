package basic

import (
	"go-microservice/common/basic/config"
	"go-microservice/common/basic/mysql"
	"go-microservice/common/basic/redis"
)

func Init() {
	config.Init()
	mysql.Init()
	redis.Init()
}
