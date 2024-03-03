package main

import (
	"backend-github-trending/db"
	"backend-github-trending/handler"
	"github.com/labstack/echo"
)

func main() {

	sql := &db.Sql{
		Host:     "localhost",
		Port:     5433,
		Username: "postgres",
		DbName:   "golang",
	}
	sql.Connect()
	defer sql.Close()

	e := echo.New()
	e.GET("/", handler.Welcome)

	e.GET("/user/sign-in", handler.HandleSignIn)
	e.GET("/user/sign-up", handler.HandleSignup)

	e.Logger.Fatal(e.Start(":3000"))
}
