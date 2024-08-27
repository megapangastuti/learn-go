package middleware

import (
	"github.com/gin-gonic/gin"
)

func SimpleMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.WriteString("Before request\n")

		ctx.Next()

		ctx.Writer.WriteString("After request\n")
	}
}
