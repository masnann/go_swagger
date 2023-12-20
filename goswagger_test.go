package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetUsers(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, getUsers(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		var users []User
		err := json.Unmarshal(rec.Body.Bytes(), &users)
		assert.NoError(t, err)

		expectedUsers := []User{
			{Name: "John", Phone: "123456"},
			{Name: "Jane", Phone: "789012"},
		}
		assert.Equal(t, expectedUsers, users)
	}
}
