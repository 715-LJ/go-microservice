package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/logx"
	"go-microservice/common/basic/config"
	"sync"
)

var (
	client *redis.Client
	m      sync.RWMutex
	inited bool
)

// Init 初始化Redis
func Init() {
	m.Lock()
	defer m.Unlock()

	if inited {
		logx.Info("已经初始化过Redis...")
		return
	}

	redisConfig := config.GetRedisConfig()
	// 打开才加载
	if redisConfig != nil && redisConfig.GetEnabled() {

		logx.Info("初始化Redis...")
		initSingle(redisConfig)
		logx.Info("初始化Redis，检测连接...")

		pong, err := client.Ping(context.TODO()).Result()
		if err != nil {
			logx.Errorf(err.Error())
		}

		logx.Info("初始化Redis，检测连接Ping. YES ", pong)
	}
}

// GetRedis 获取redis
func GetRedis() *redis.Client {
	return client
}

func initSentinel(redisConfig config.RedisConfig, redisSentConfig config.RedisSentinelConfig) {
	client = redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:    redisSentConfig.GetMaster(),
		SentinelAddrs: redisSentConfig.GetNodes(),
		DB:            redisConfig.GetDBNum(),
		Password:      redisConfig.GetPassword(),
	})

}

func initSingle(redisConfig config.RedisConfig) {
	client = redis.NewClient(&redis.Options{
		Addr:     redisConfig.GetConn(),
		Password: redisConfig.GetPassword(), // no password set
		DB:       redisConfig.GetDBNum(),    // use default DB
	})
}
