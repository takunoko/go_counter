package main

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/redis/go-redis/v9"
	myRedis "go_web_counter/infrastructure/redis"
	"net/http"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", hello)
	// Routes
	e.GET("/redis_test", redisHandler)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, New-3 World!")
}

// Redis Hnadler
func redisHandler(echoCtx echo.Context) error {
	ctx := context.Background()

	key1 := "k"

	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	repo := myRedis.NewDataRepository(rdb)

	retStr := ""

	val := 111
	err := repo.Set(ctx, key1, val)
	if err != nil {
		panic(err)
	}
	retStr += fmt.Sprintf("Set '%v': %v\n", key1, val)

	val, err = repo.Get(ctx, key1)
	if err != nil {
		panic(err)
	}
	retStr += fmt.Sprintf("Get '%v': %v\n", key1, val)

	val, err = repo.CntUp(ctx, key1)
	if err != nil {
		panic(err)
	}
	retStr += fmt.Sprintf("CntUp '%v': %v\n", key1, val)

	echoCtx.String(http.StatusOK, retStr)

	return nil
}
