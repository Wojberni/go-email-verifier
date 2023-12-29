package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	_ "go-email-verifier/docs"
	user "go-email-verifier/user"
)

func main() {
	e := echo.New()

	// Middleware configuration
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Allows requests from any origin wth GET, HEAD, PUT, POST or DELETE method.
	e.Use(middleware.CORS())

	// CORS restricted
	// Allows requests from any `https://labstack.com` or `https://labstack.net` origin
	// wth GET, PUT, POST or DELETE method.
	// e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	// 	AllowOrigins: []string{"https://labstack.com", "https://labstack.net"},
	// 	AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	// }))

	// Routes configuration
	g := e.Group("/v2/users")
	g.POST("", user.CreateUser)
	g.GET("", user.GetAllUsers)
	g.GET("/:id", user.GetUser)
	g.PUT("/:id", user.UpdateUser)
	g.DELETE("/:id", user.DeleteUser)

	e.Logger.Fatal(e.Start(":1323"))
}
