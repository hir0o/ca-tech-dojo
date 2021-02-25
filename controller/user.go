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
func (connect *ConnectDB) UserCreate(c echo.Context) error {
	// TODO: newと := の違い
	user := new(NameJson) // jsonの受け取り
	if err := c.Bind(user); err != nil {
		return err
	}

	token, err := record.CreateUser(user.Name, connect.DB) // userの作成

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	res := TokenJson{
		Token: token,
	}
	return c.JSON(http.StatusOK, res)
}

// UserGet GET /user/get
func (connect *ConnectDB)UserGet(c echo.Context) error {
	// x-tokenの取得
	token := c.Request().Header.Get("x-token")
	user, err := record.GetUser(token, connect.DB)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	res := NameJson{
		Name: user.Name,
	}

	return c.JSON(http.StatusOK, res)
}

// UserUpdate PUT /user/update
func (connect *ConnectDB) UserUpdate(c echo.Context) error {
	user := new(NameJson)
	if err := c.Bind(user); err != nil {
		return err
	}
	// x-tokenの取得
	token := c.Request().Header.Get("x-token")

	if err := record.UpdateUser(user.Name, token, connect.DB); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	res := NameJson{
		Name: user.Name,
	}

	return c.JSON(http.StatusOK, res)
}
