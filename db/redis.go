package db

import (
	"fmt"
	"github.com/go-redis/redis"
	"go-gin-demo/config"
)

var RedisClient *redis.Client

func InitRedis(cfg *config.Config) {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Redis.Host, cfg.Redis.Port),
		Password: cfg.Redis.Auth, // no password set
		DB:       cfg.Redis.Db,   // use default DB
	})

	_, err := RedisClient.Ping().Result()
	if err != nil {
		panic("redis connect error")
	}
}
