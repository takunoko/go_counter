package test

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"os"
	"testing"
)

func OpenTestRedis(t *testing.T) *redis.Client {
	t.Parallel()

	redisHost := os.Getenv("TEST_REDIS_HOST")
	if redisHost == "" {
		redisHost = "127.0.0.1" // Default
	}
	redisPort := os.Getenv("TEST_REDIS_PORT")
	if redisPort == "" {
		redisHost = "6379" // Default
	}
	fmt.Printf("Host: %s, Port: %s", redisHost, redisPort)

	testRCli := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", redisHost, redisPort),
		Password: "",
		DB:       0,
	})

	// データの初期化
	ctx := context.Background()
	testRCli.FlushDBAsync(ctx)

	return testRCli
}
