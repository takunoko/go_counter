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

	err := rdb.Set(ctx, "k", "123", 0).Err()
	if err != nil {
		panic(err)
	}

	retStr := "set key: k, val: 123\n"

	val, err := rdb.Get(ctx, "k").Result()
	if err != nil {
		panic(err)
	}
	retStr += fmt.Sprintf("get key: k, val: %s\n", val)

	val = strconv.FormatInt(rdb.Incr(ctx, "key").Val(), 10)
	retStr += fmt.Sprintf("incr key: k, val: %s\n", val)

	val2, err := rdb.Get(ctx, "key2").Result()
	if err == redis.Nil {
		retStr += "key: key2 is not found\n"
	} else if err != nil {
		panic(err)
	} else {
		retStr += fmt.Sprintf("incr key2: k, val: %s\n", val2)
	}

	echoCtx.String(http.StatusOK, retStr)

	// Output: key value
	// key2 does not exist
	return nil
}
