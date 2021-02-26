package controller

import (
	"ca-tech-dojo/record"
	"net/http"

	"github.com/labstack/echo"
)

type TimesJson struct {
	Times int `json:"times"`
}

type ResultJson struct {
	Results []record.GachaResult `json:"results"`
}

// GachaDraw  /gacha/draw
func  (connect *ConnectDB)GachaDraw(c echo.Context) (err error) {
	times := new(TimesJson)
	if err := c.Bind(times); err != nil {
		return err
	}
	token := c.Request().Header.Get("x-token")

	characters, err := record.GachaDraw(times.Times, token, connect.DB)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	res := ResultJson{
		Results: characters,
	}
	return c.JSON(http.StatusOK, res)
}
