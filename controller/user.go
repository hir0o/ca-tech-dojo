package controller

import (
	"ca-tech-dojo/record"
	"net/http"

	"github.com/labstack/echo"
)

type reqUserCreate struct {
	Name string `json:"name"`
}

type resUserCreate struct {
	Token string `json:"token"`
}

// POST /user/create
func UserCreate(c echo.Context) (err error) {
	user := new(reqUserCreate) // jsonの受け取り
	if err := c.Bind(user); err != nil {
		return err
	}

	token := record.CreateUser(user.Name) // userの作成
	res := resUserCreate{
		Token: token,
	}
	return c.JSON(http.StatusOK, res)
}

type resUserGet struct {
	Name string `json:"name"`
}

// GET /user/get
func UserGet(c echo.Context) (err error) {
	// x-tokenの取得
	xTokne := c.Request().Header.Get("x-token")
	username := record.GetUser(xTokne)

	res := resUserGet{
		Name: username,
	}
	return c.JSON(http.StatusOK, res)
}
