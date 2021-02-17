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
	token := c.Request().Header.Get("x-token") // タイポ

	characters := record.CharacterList(token)

	res := CharactersJson{
		Characters: characters,
	}

	return c.JSON(http.StatusOK, res)
}
