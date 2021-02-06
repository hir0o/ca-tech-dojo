package controller

import (
	"ca-tech-dojo/record"
	"net/http"

	"github.com/labstack/echo"
)

type NameJson struct {
	Name string `json:"name"`
}

type TokenJson struct {
	Token string `json:"token"`
}

// UserCreate POST /user/create
func UserCreate(c echo.Context) (err error) {
	user := new(NameJson) // jsonの受け取り
	if err := c.Bind(user); err != nil {
		return err
	}

	token := record.CreateUser(user.Name) // userの作成
	res := TokenJson{
		Token: token,
	}
	return c.JSON(http.StatusOK, res)
}

// UserGet GET /user/get
func UserGet(c echo.Context) (err error) {
	// x-tokenの取得
	xTokne := c.Request().Header.Get("x-token")
	username, err := record.GetUser(xTokne)

	// errがあったら、403を返す
	if err != nil {
		return c.NoContent(http.StatusForbidden)
	}

	res := NameJson{
		Name: username,
	}
	return c.JSON(http.StatusOK, res)
}

// UserUpdate PUT /user/update
func UserUpdate(c echo.Context) (err error) {
	user := new(NameJson)
	if err := c.Bind(user); err != nil {
		return err
	}
	// x-tokenの取得
	xTokne := c.Request().Header.Get("x-token")

	if err := record.UpdateUser(user.Name, xTokne); err != nil {
		// errがあったら、403
		return c.NoContent(http.StatusForbidden)
	}

	res := NameJson{
		Name: user.Name,
	}

	return c.JSON(http.StatusOK, res)
}
