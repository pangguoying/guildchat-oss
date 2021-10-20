package redis

import (

	"github.com/go-redis/redis"
	"guildchat-oss/configs"
)

var RedisClient *redis.Client

func init() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     configs.App.RedisConf.Addr,
		Password: configs.App.RedisConf.Password,
		DB:       configs.App.RedisConf.Db,
	})

	err := RedisClient.Ping().Err()
	if err != nil {
		panic("redis connect error")
	}
}
