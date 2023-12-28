package main

import (
	"sync"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	users = map[int]*User{}
	seq   = 1
	lock  = sync.Mutex{}
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/v2/users", createUser)
	e.GET("/v2/users", getAllUsers)

	e.Logger.Fatal(e.Start(":1323"))
}
