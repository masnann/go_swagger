package main

import (
	"github.com/labstack/echo/v4"
	"github.com/swaggo/echo-swagger"
	_ "go_swagger/docs"
)

type User struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

func main() {
	e := echo.New()

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.GET("/users", getUsers)

	e.Start(":8080")
}
