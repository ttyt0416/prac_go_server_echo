package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/ttyt0416/learngo/configs"
	"github.com/ttyt0416/learngo/post"
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
	configs.ConnectDB()

	//user methods
	e.GET("/users", user.GetAllUsers)
	e.GET("/users/:userId", user.GetUser)
	e.PUT("/users/:userId", user.UpdateUser)
	e.POST("/users", user.CreateUser)
	e.DELETE("/users/:userId", user.DeleteUser)
	e.POST("/login", user.Login)

	//post methods
	e.GET("/posts", post.GetAllPosts)
	e.GET("/posts/:postId", post.GetPost)
	e.PUT("/posts/:postId", post.UpdatePost)
	e.POST("/posts", post.CreatePost)
	e.DELETE("/posts/:postId", post.DeletePost)

	// Unauthenticated route
	e.GET("/", user.Accessible)
	// Restricted group
	r := e.Group("/restricted")

	// Configure middleware with the custom claims type
	config := middleware.JWTConfig{
		Claims:     &user.JwtCustomClaims{},
		SigningKey: []byte("secret"),
	}
	r.Use(middleware.JWTWithConfig(config))
	r.GET("", user.Restricted)

	e.Logger.Fatal(e.Start(":1323"))
}
