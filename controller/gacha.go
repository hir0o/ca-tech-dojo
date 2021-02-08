package controller

import (
	"ca-tech-dojo/record"
	"net/http"

	"github.com/labstack/echo"
)

type TimesJson struct {
	Times int `json:"name"`
}

type ResultJson struct {
	Results []record.Charactor
}

type Charactor struct {
	ID            int
	CharactorRank int
	Name          string
}

// GachaDraw  /gacha/draw
func GachaDraw(c echo.Context) (err error) {
	times := new(TimesJson)
	xTokne := c.Request().Header.Get("x-token")
	if err := c.Bind(times); err != nil {
		return err
	}

	charactors := record.GachaDraw(times.Times, xTokne)

	res := ResultJson{
		Results: charactors,
	}
	return c.JSON(http.StatusOK, res)
}
