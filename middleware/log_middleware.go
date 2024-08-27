package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func LogMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		t := time.Now()

		ctx.Next()

		// Kita atur setelah akses endpoint berhasil keterangannya mau apa?
		/*
			1. Waktu akses endpoint (latency)
			2. IP Client
			3. Method [GET, POST, DELETE, PUT, ...]
			4. Status Code [1xx, 2xx, 3xx, 4xx, 5xx]
			5. UserAgent() => [Postman, Chrome, Mozzila, Safari ...]
			6. Path [/users, /products, ....]
		*/
		latency := time.Since(t)
		clientIP := ctx.ClientIP() // 192.168.1.1
		method := ctx.Request.Method
		statusCode := ctx.Writer.Status()
		userAgent := ctx.Request.UserAgent()
		path := ctx.Request.URL.Path
		log.Printf("[LOG] %s - [%v] \"%s %s %d %v \"%s\"\n",
			clientIP,
			t,
			method,
			path,
			statusCode,
			latency,
			userAgent)
	}
}
