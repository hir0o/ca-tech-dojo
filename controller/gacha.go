package controller

import (
	"ca-tech-dojo/record"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type TimesJson struct {
	Times string `json:"times"`
}

type ResultJson struct {
	Results []record.GachaResult `json:"results"`
}

// GachaDraw  /gacha/draw
func GachaDraw(c echo.Context) (err error) {
	times := new(TimesJson)
	if err := c.Bind(times); err != nil {
		return err
	}

	// 型のキャスト
	timesInt, _ := strconv.Atoi(times.Times)
	// x-tokenを受け取る
	xTokne := c.Request().Header.Get("x-token")

	characters := record.GachaDraw(timesInt, xTokne)

	res := ResultJson{
		Results: characters,
	}
	return c.JSON(http.StatusOK, res)
}
