package app

import (
	"context"
	"fmt"
	"portfolio/services/infrastructure/config"

	"github.com/go-redis/redis/v8"
)

func initRDB() (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d",
			config.StringWithDefault("redis.host", "localhost"),
			config.IntWithDefault("redis.port", 6379),
		),
		Password: config.StringWithDefault("redis.password", ""),
		DB:       config.IntWithDefault("redis.db", 0),
	})
	if err := rdb.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}
	return rdb, nil
}
