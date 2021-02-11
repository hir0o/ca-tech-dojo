package controller

import (
	"ca-tech-dojo/record"
	"net/http"

	"github.com/labstack/echo"
)

type CharactersJson struct {
	Characters []record.Character `json:"characters"`
}

func CharacterList(c echo.Context) (err error) {
	xTokne := c.Request().Header.Get("x-token")

	characters := record.CharacterList(xTokne)

	res := CharactersJson{
		Characters: characters,
	}

	return c.JSON(http.StatusOK, res)
}
