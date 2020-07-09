package server

import (
	"github.com/go-redis/redis"
)

var Redis *redis.Client

func InitRedis() {
	Redis = redis.NewClient(&redis.Options{
		Addr: "redis:6379",
	})
}
