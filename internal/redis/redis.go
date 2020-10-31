package redis

import (
	"fmt"
	"os"

	"github.com/go-redis/redis"
)

var Redis *redis.Client

func InitRedis() {
	addr := fmt.Sprintf("%s:6379", os.Getenv("REDIS_HOST"))
	Redis = redis.NewClient(&redis.Options{
		Addr: addr,
	})
}
