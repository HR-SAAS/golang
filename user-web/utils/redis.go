package utils

import (
	"fmt"
	redis "github.com/go-redis/redis/v8"
	"hr-saas-go/user-web/global"
)

func NewRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", global.Config.RedisConfig.Host, global.Config.RedisConfig.Port),
	})
	return rdb
}
