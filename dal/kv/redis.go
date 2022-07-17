package kv

import (
	"errors"

	"github.com/go-redis/redis/v8"
	"github.com/mc_nft/conf"
)

var RedisClient *redis.Client

func InitRedis() {
	redisConf := conf.GetInstance().Redis
	if redisConf.Host == "" {
		panic(errors.New("invalid redis conf"))
	}
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     redisConf.Host,
		Password: redisConf.Password,
		DB:       redisConf.DB,
	})
}
