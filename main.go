package main

import (
	"ca-tech-dojo/controller"
	"ca-tech-dojo/db"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	db.Init()
	e.POST("/user/create", controller.UserCreate)
	e.GET("/user/get", controller.UserGet)
	e.PUT("/user/update", controller.UserUpdate)
	e.Logger.Fatal(e.Start(":8080"))
}
