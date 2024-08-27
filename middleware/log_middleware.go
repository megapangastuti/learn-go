package middleware

import (
	"incubation/model"
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

		model.SendLogRequest(model.LogModel{
			AccessTime: t,
			Latency:    time.Since(t),
			ClientIP:   ctx.ClientIP(), // 192.168.1.1
			Method:     ctx.Request.Method,
			Code:       ctx.Writer.Status(),
			Path:       ctx.Request.URL.Path,
			UserAgent:  ctx.Request.UserAgent(),
		})
	}
}
