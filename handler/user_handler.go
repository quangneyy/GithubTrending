package handler

import (
	"github.com/labstack/echo"
	"net/http"
)

func HandleSignIn(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"user":  "Quang",
		"email": "quang@gmail.com",
	})
}

func HandleSignup(c echo.Context) error {
	type User struct {
		Email    string `json:"email"`
		FullName string `json:"name"`
		Age      int    `json:"age"`
	}

	user := User{
		Email:    "quang@gmail.com",
		FullName: "quang",
		Age:      90,
	}
	return c.JSON(http.StatusOK, user)
}
