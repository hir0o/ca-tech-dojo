package controller

import (
	"net/http"

	"github.com/labstack/echo"
)

type reqUserCreate struct {
	Name string `json:"name"`
}

// POST /user/create
func UserCreate(c echo.Context) (err error) {
	// x-tokenの取得
	xTokne := c.Request().Header.Get("x-token")
	println(xTokne)
	user := new(reqUserCreate)
	if err := c.Bind(user); err != nil {
		return err
	}
	// とりあえずそのまま返す
	return c.JSON(http.StatusOK, user)
}

type resUserGet struct {
	Name string `json:"name"`
}

// GET /user/get
func UserGet(c echo.Context) (err error) {
	res := resUserGet{
		Name: "username",
	}
	return c.JSON(http.StatusOK, res)
}
