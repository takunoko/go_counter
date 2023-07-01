package web

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (a ApiHandlers) GetNum(ctx echo.Context) error {
	var res GetNumResponse
	if err := Initialize(&res.JSON200); err != nil {
		panic(err)
	}

	currentNum := CurrentNum{
		Num: 8,
	}

	res.JSON200.Result = currentNum

	return ctx.JSON(http.StatusOK, res.JSON200)
}
