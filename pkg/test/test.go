package test

import (
	"context"
	"github.com/redis/go-redis/v9"
	"go_web_counter/config"
	"testing"
)

const (
	testRedisAddress = "127.0.0.1:6379" // テストのGoはローカルで実行なためIPで指定
	testRedisPass    = config.RedisPass
	testRedisDBNo    = 9 // テスト用で分ける
)

func OpenTestRedis(t *testing.T) *redis.Client {
	t.Parallel()

	testRCli := redis.NewClient(&redis.Options{
		Addr:     testRedisAddress,
		Password: testRedisPass,
		DB:       testRedisDBNo,
	})

	// データの初期化
	ctx := context.Background()
	testRCli.FlushDBAsync(ctx)

	return testRCli
}
