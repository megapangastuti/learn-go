package main

import (
	"incubation/handler"
	"incubation/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	// engine := gin.Default()
	engine := gin.New()

	// engine.Use(middleware.SimpleMiddleware())
	// Menggunakan middleware secara global route
	engine.Use(middleware.LogMiddleware())

	// Selanjutnya kita bisa langsung menggunakan fiturnya :
	// 1. Menggunakan method => GET, POST, PUT, DELETE

	// Menerima sebuah parameter [1] REST URL, [Handler]
	// http://enigmacamp.com:8888/users
	// RelativePath => Noun (Kata benda) (Plural)

	// GET
	// Menggunakan middleware secara group route
	v1 := engine.Group("/api/v1", middleware.BasicAuthMiddleware())
	v1.GET("/ping", handler.PingHandler)
	v1.GET("/pong", handler.PongHandler)
	v1.GET("/greetings", handler.GreetingsHandler) // karena opsional, jadi ditaro diatas yang required. kalo
	// engine.GET("/greetings/:name", handler.GreetingByNameHandler)
	v1.GET("/greetings/:params", handler.GreetingByOtherHandler)

	//POST
	v1_Post := engine.Group("/api/v1", middleware.LogMiddleware())
	v1_Post.POST("/user", handler.CreateUCHandler)
	v1_Post.POST("/user/register", handler.CreateUserCredentialWithPhotoHandler)

	engine.Run()
}
