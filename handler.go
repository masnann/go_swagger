package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

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
