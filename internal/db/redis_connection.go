package db

import (
	"context"
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func RedisConnect() *redis.Client {
	host := os.Getenv("REDIS_HOST")
	port := os.Getenv("REDIS_PORT")
	password := os.Getenv("REDIS_PASSWORD")

	addr := fmt.Sprintf("%s:%s", host, port)

	rdb := redis.NewClient((&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       0,
	}))

	err := rdb.Ping(ctx).Err()
	if err != nil {
		panic(fmt.Sprintf("failed to connect to Redis: %v", err))
	}
	return rdb
}
