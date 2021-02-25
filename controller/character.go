package controller

import (
	"ca-tech-dojo/record"
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo"
)

type CharactersJson struct {
	Characters []record.Character `json:"characters"`
}

func (connect *ConnectDB) CharacterList(c echo.Context) (err error) {
	token := c.Request().Header.Get("x-token")

	characters, err := record.CharacterList(token, connect.DB)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	res := CharactersJson{
		Characters: characters,
	}

	return c.JSON(http.StatusOK, res)
}
