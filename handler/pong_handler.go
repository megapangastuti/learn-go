package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PongHandler(ctx *gin.Context) {
	ctx.String(http.StatusOK, "pong")
}
