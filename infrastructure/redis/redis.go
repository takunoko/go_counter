package myRedis

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
	"net/http"
	"strconv"
)

var ctx = context.Background()

func RedisClientTest(echoCtx echo.Context) error {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	key1 := "k"
	key2 := "k2"

	err := rdb.Set(ctx, key1, "123", 0).Err()
	if err != nil {
		panic(err)
	}

	retStr := fmt.Sprintf("set key: %v, val: 123\n", key1)

	val, err := rdb.Get(ctx, key1).Result()
	if err != nil {
		panic(err)
	}
	retStr += fmt.Sprintf("get key: %v, val: %s\n", key1, val)

	incrVal := strconv.FormatInt(rdb.Incr(ctx, key1).Val(), 10)
	retStr += fmt.Sprintf("incr key: %v, val: %s\n", key1, incrVal)

	val2, err := rdb.Get(ctx, key2).Result()
	if err == redis.Nil {
		retStr += fmt.Sprintf("key: %v is not found\n", key2)
	} else if err != nil {
		panic(err)
	} else {
		retStr += fmt.Sprintf("incr %v, val: %s\n", key2, val2)
	}

	echoCtx.String(http.StatusOK, retStr)

	// Output: key value
	// key2 does not exist
	return nil
}
