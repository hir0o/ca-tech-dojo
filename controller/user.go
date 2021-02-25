package controller

import (
	"ca-tech-dojo/record"
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo"
)

type NameJson struct {
	Name string `json:"name"`
}

type TokenJson struct {
	Token string `json:"token"`
}

// UserCreate POST /user/create
func (connect *Connect) UserCreate(c echo.Context) (err error) {
	// TODO: newと := の違い
	user := new(NameJson) // jsonの受け取り
	if err := c.Bind(user); err != nil {
		return err
	}

	token, err := record.CreateUser(user.Name, connect.DB) // userの作成

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	res := TokenJson{
		Token: token,
	}

	return c.JSON(http.StatusOK, res)
}

// UserGet GET /user/get
func (connect *Connect) UserGet(c echo.Context) (err error) {
	// x-tokenの取得
	token := c.Request().Header.Get("x-token")
	user, err := record.GetUser(token, connect.DB) // user形で返す

	// errがあったら、403を返す 500エラーも返す
	if err != nil {
		return c.NoContent(http.StatusForbidden)
	}
	res := NameJson{
		Name: username,
	}

	return c.JSON(http.StatusOK, res)
}

// UserUpdate PUT /user/update
func (connect *Connect) UserUpdate(c echo.Context) (err error) {
	user := new(NameJson)
	if err := c.Bind(user); err != nil {
		return err
	}
	// x-tokenの取得
	token := c.Request().Header.Get("x-token")

	if err := record.UpdateUser(user.Name, token, connect.DB); err != nil {
		// errがあったら、403
		return c.NoContent(http.StatusForbidden)
	}

	res := NameJson{
		Name: user.Name,
	}

	return c.JSON(http.StatusOK, res)
}
