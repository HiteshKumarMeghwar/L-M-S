package database

import (
	"fmt"

	"github.com/HiteshKumarMeghwar/L-M-S/config"
	"github.com/go-redis/redis/v8"
)

func ConnectionRedisDb(config *config.Config) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: config.RedisUrl,
	})

	fmt.Println("Connection successfully ... (redis)")

	return rdb
}
