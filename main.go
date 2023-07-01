package main

import (
	"context"
	"fmt"
	oapiMiddleware "github.com/deepmap/oapi-codegen/pkg/middleware"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/redis/go-redis/v9"
	"go_web_counter/config"
	myRedis "go_web_counter/infrastructure/redis"
	"go_web_counter/interface/web"
	"net/http"
)

type apiController struct{}

func (a apiController) GetNum(ctx echo.Context) error {
	currentNum := web.CurrentNum{
		Num: 8,
	}
	json200 := struct {
		Result web.CurrentNum `json:"result"`
	}{
		Result: currentNum,
	}
	var res web.GetNumResponse
	// if err := defaults.Set(&res.JSON200); err != nil {
	// 	panic(err)
	// }
	res.JSON200 = &json200

	return ctx.JSON(http.StatusOK, res.JSON200)
}

func main() {
	// Echo instance
	e := echo.New()

	// OpenApiの定義に従ったリクエストかのバリデーションミドルウェア
	swagger, err := web.GetSwagger()
	if err != nil {
		panic(err)
	}
	e.Use(oapiMiddleware.OapiRequestValidator(swagger))

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//// Routes
	//e.GET("/hello", helloHandler)
	//e.GET("/", getHandler)
	//e.GET("/reset", resetHandler)
	//e.GET("/cnt_up", cntUpHandler)
	//e.GET("/cnt_down", cntDownHandler)

	api := apiController{}
	web.RegisterHandlers(e, api)

	// Start server
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", config.ServerPort)))
}

// Handler
func helloHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

const (
	Key = "cnt_key"
)

// TODO: DIしたい
func getRedisCli() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     config.RedisAddr,
		Password: config.RedisPass,
		DB:       config.RedisDB,
	})
}

func getHandler(echoCtx echo.Context) error {
	ctx := context.Background()
	rCli := getRedisCli()
	repo := myRedis.NewDataRepository(rCli)

	val, err := repo.Get(ctx, Key)
	if err != nil {
		panic(err)
	}

	dispStr := fmtDisp("CntGet", Key, val)

	echoCtx.String(http.StatusOK, dispStr)

	return nil
}

func resetHandler(echoCtx echo.Context) error {
	ctx := context.Background()
	rCli := getRedisCli()
	repo := myRedis.NewDataRepository(rCli)

	initialVal := 0
	err := repo.Set(ctx, Key, initialVal)
	if err != nil {
		panic(err)
	}

	dispStr := fmtDisp("Set", Key, initialVal)

	echoCtx.String(http.StatusOK, dispStr)

	return nil
}

func cntUpHandler(echoCtx echo.Context) error {
	ctx := context.Background()
	rCli := getRedisCli()
	repo := myRedis.NewDataRepository(rCli)

	val, err := repo.CntUp(ctx, Key)
	if err != nil {
		panic(err)
	}

	dispStr := fmtDisp("CntUp", Key, val)

	echoCtx.String(http.StatusOK, dispStr)

	return nil
}

func cntDownHandler(echoCtx echo.Context) error {
	ctx := context.Background()
	rCli := getRedisCli()
	repo := myRedis.NewDataRepository(rCli)

	val, err := repo.CntDown(ctx, Key)
	if err != nil {
		panic(err)
	}

	dispStr := fmtDisp("CntDown", Key, val)

	echoCtx.String(http.StatusOK, dispStr)

	return nil
}

func fmtDisp(f string, key string, v int) string {
	return fmt.Sprintf("%10s '%v': %3d\n", f, key, v)
}
