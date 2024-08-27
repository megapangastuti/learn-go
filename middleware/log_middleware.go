package middleware

import (
	"incubation/model"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func LogMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// os Package OpenFile, fungsinya untuk membuka akses ke file tsb
		// logger.log (file yang akan diberikan akses)
		// APPEND => akan menambahkan data di akhir (append slice)
		// CREATE => akan membuat file logger.log jika tidak ada
		// WRONLY =? akan memberikan akses menulis ke file (isi ke file)
		// ModePerm => Permission atau bisa ganti dengan 0644
		file, err := os.OpenFile("logger.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModePerm)

		if err != nil {
			log.Fatal("err:", err.Error())
		}

		defer file.Close()
		log.SetOutput(file)

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

		logString := model.SendLogRequest(model.LogModel{
			AccessTime: t,
			Latency:    time.Since(t),
			ClientIP:   ctx.ClientIP(), // 192.168.1.1
			Method:     ctx.Request.Method,
			Code:       ctx.Writer.Status(),
			Path:       ctx.Request.URL.Path,
			UserAgent:  ctx.Request.UserAgent(),
		})

		// untuk menulis ke file yang sudah dibuka dan diberi akses
		_, err = file.WriteString(logString)
		if err != nil {
			log.Fatal("faile to write", err.Error())
		}
	}
}
