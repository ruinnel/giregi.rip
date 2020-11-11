package common

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

func OpenRedis(config Redis) *redis.Client {
	logger := GetLogger()
	ctx := context.Background()
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.Host, config.Port),
		Password: config.Password,
		DB:       config.Db,
	})
	_, err := client.Ping(ctx).Result()
	if err != nil {
		logger.Panicf("failed to connect redis - %v", config)
	}
	return client
}
