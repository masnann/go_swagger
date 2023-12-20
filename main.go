package main

import (
	"github.com/labstack/echo/v4"
	"github.com/swaggo/echo-swagger"
	_ "go_swagger/docs"
	"net/http"
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

// @Summary Get list of users
// @Description Get all users
// @Produce json
// @Success 200 {array} User
// @Router /users [get]
func getUsers(c echo.Context) error {
	// Logic untuk mendapatkan data users
	users := []User{
		{Name: "John", Phone: "123456"},
		{Name: "Jane", Phone: "789012"},
	}

	return c.JSON(http.StatusOK, users)
}
