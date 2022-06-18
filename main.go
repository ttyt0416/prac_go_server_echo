package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/ttyt0416/learngo/auth"
	"github.com/ttyt0416/learngo/user"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	//user methods
	e.GET("/users", user.GetAllUsers)
	e.GET("/users/:id", user.GetUser)
	e.PUT("/users/:id", user.UpdateUser)
	e.POST("/users", user.CreateUser)
	e.DELETE("/users/:id", user.DeleteUser)

	//auth methods

	// Login route
	e.POST("/login", auth.Login)
	// Unauthenticated route
	e.GET("/", auth.Accessible)
	// Restricted group
	r := e.Group("/restricted")

	// Configure middleware with the custom claims type
	config := middleware.JWTConfig{
		Claims:     &auth.JwtCustomClaims{},
		SigningKey: []byte("secret"),
	}
	r.Use(middleware.JWTWithConfig(config))
	r.GET("", auth.Restricted)

	e.Logger.Fatal(e.Start(":1323"))
}
