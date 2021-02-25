package main

import (
	"ca-tech-dojo/controller"
	"ca-tech-dojo/db"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	db := db.Init()

	handler := controller.Connect{
    DB: db,
  }
	e.POST("/user/create", handler.UserCreate)
	e.GET("/user/get", handler.UserGet)
	e.PUT("/user/update", handler.UserUpdate)
	e.POST("/gacha/draw", handler.GachaDraw)
	e.GET("/character/list", handler.CharacterList)
	e.Logger.Fatal(e.Start(":8080"))
}
