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
	engine.Use(middleware.LogMiddleware())

	engine.GET("/ping", handler.PingHandler)

	engine.GET("/pong", handler.PongHandler)
	engine.GET("/greetings", handler.GreetingsHandler) // karena opsional, jadi ditaro diatas yang required. kalo
	// engine.GET("/greetings/:name", handler.GreetingByNameHandler)
	engine.GET("/greetings/:params", handler.GreetingByOtherHandler)

	engine.POST("/user", handler.CreateUCHandler)
	engine.POST("/user/register", handler.CreateUserCredentialWithPhotoHandler)

	engine.Run()
}
