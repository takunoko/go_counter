package test

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

const (
	testRedisHost = "127.0.0.1"
	testRedisPort = "6379"
	testRedisDB   = 9
)

func OpenTestRedis() *redis.Client {
	testRCli := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", testRedisHost, testRedisPort),
		Password: "",
		DB:       testRedisDB,
	})

	// データの初期化
	ctx := context.Background()
	testRCli.FlushDBAsync(ctx)

	return testRCli
}
