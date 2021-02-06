package main

import (
	"ca-tech-dojo/controller"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.POST("/user/create", controller.UserCreate)
	e.GET("/user/get", controller.UserGet)
}
