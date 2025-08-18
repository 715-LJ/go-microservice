package config

import (
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"path/filepath"
	"sync"
)

var (
	mysqlConfig defaultMysqlConfig
	redisConfig defaultRedisConfig
	m           sync.RWMutex
	inited      bool
)

// Init 初始化配置
func Init() {

	m.Lock()
	defer m.Unlock()

	if inited {
		logx.Info("[Init] 配置已经初始化过")
		return
	}

	// 加载yml配置
	conf.MustLoad(filepath.Join("./", "etc", "mysql.yaml"), &mysqlConfig)
	conf.MustLoad(filepath.Join("./", "etc", "redis.yaml"), &redisConfig)

	// 标记已经初始化
	inited = true
}

func GetMysqlConfig() (ret MysqlConfig) {
	return mysqlConfig
}

// GetRedisConfig 获取redis配置
func GetRedisConfig() (ret RedisConfig) {
	return redisConfig
}
