package handler

import (
	"fmt"
	"incubation/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GreetingsHandler(ctx *gin.Context) {
	// menggunakan query param
	// URL => greetings?name=Mega&address=Solo
	// di Gin itu sudah disediakan dengan nama .Query => ctx.Query("name") && ctx.Query("address")
	name := ctx.DefaultQuery("name", "Mega")       // optional dikirimkan di url path
	address := ctx.DefaultQuery("address", "Solo") // optional dikirimkan di url path
	// resp := fmt.Sprintf("Hai ... %s yang lagi ada di %s", name, address)

	// gin.H => sebenernya adalah map[string]interface{}
	// "key": "value" | 123 | true
	// ctx.JSON(http.StatusOK, gin.H{
	// 	"code":    http.StatusOK,
	// 	"message": "Ok",
	// 	"data":    resp,
	// })

	// memulai menggunakan model response
	resp := fmt.Sprintf("Hai ... %s yang lagi ada di %s", name, address)
	model.SendSingleResponse(ctx, "Okee", resp)

}

func GreetingByNameHandler(ctx *gin.Context) {
	// menggunakan sebuah parameter
	// URL => greetings/:name => greetings/Mega Pangastuti
	// di Gin itu sudah disediakan dengan nama .Query => ctx.Param("name")
	name := ctx.Param("name") // required
	ctx.String(http.StatusOK, "Hai ... %s", name)
}

func GreetingByOtherHandler(ctx *gin.Context) {
	params := ctx.Params
	ctx.String(http.StatusOK, "%v", params)
}
