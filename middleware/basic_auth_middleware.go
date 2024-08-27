package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func BasicAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// Di Gin sudah menyediakan yang namanya BasicAuth()
		username, password, ok := ctx.Request.BasicAuth()
		fmt.Println("username : ", username)
		fmt.Println("password : ", password)
		fmt.Println("ok : ", ok)

		// Logic pengecekan username dan password
		if !ok || username != "admin" || password != "password" {
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}

		ctx.Next()
	}
}
