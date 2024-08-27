package handler

import (
	"encoding/json"
	"fmt"
	"incubation/model"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

var UCDb []model.UserCredential

func CreateUCHandler(ctx *gin.Context) {
	/*
		TODO :
		1. kita siapkan sebuah payload (struct UserCredential)
		2. Kemudian kita validasi menggunakan :
		   - ctx.MustBind
		   - ctx.MustBindJSON
		   - ctx.ShouldBind
		   - ctx.ShouldBindJSON
		   - dll
		3. Setelah berhasil divalidasi, kita bisa simpan hasilnya, mekanisme nya adalah :
		   - simpan di slice saja => []UserCredential
		4. Kita  balikan response hasil yang diinput/request body nya
	*/

	// Siapkan payload
	var payload model.UserCredential

	// Validasi payload
	err := ctx.ShouldBind(&payload)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// Simpan ke slice
	UCDb = append(UCDb, payload)

	// Balikin ke responnya
	model.HasBeedCreated(ctx, "berhasil menambahkan data user baru", payload)
}

// walaupun update sebenarnya insert
func CreateUserCredentialWithPhotoHandler(ctx *gin.Context) {
	//TODO :
	/*
		1. Kita melakukan insert data user sekaligus upload photo
		2. Body payload nya akan berbentuk form-data
		3. Kemudian data email dan photo akan tersimpan di payload
		4. Photo hanya akan mengambil pathnya saja
	*/

	// Kita harus tangkap json user, kita bisa memanfaatkan context gin
	// ctx.PostForm("key")
	user := ctx.PostForm("user")

	// Kita ambil filenya, denga cara :
	// ctx. Request.FormFile("key")
	file, header, err := ctx.Request.FormFile("photo")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	defer file.Close()

	// Kita siapkan sebuah path dimana photo itu disimpan
	// uploads/gambar.jpg => save di database
	// Go => file => kita bisa membuat dan menulis file
	// kita juga bisa menggabungkan nama path + nama filenya
	// Nama file bisa di ubah yaa
	// bisa menggunakan random string atau pattern => username_photo.jpg
	newFileName := fmt.Sprintf("%v_photo%s", rand.New(rand.NewSource(time.Now().UTC().UnixNano())).Int(), filepath.Ext(header.Filename))
	// fileLocation := filepath.Join("uploads", header.Filename) // uploads/gambar.jpg
	fileLocation := filepath.Join("uploads", newFileName)

	// setelah itu kita hatus buat folder nya untuk menyimpan gambar salinannya
	// buat folder secara otomatis
	// os package
	os.MkdirAll("uploads", os.ModePerm) // buat sekaligus kasih permission 0666

	// // setelah itu kita buat filenya yang diambil dari filelocation
	// // os package
	// outFile, _ := os.Create(fileLocation)

	// // salin gambar yang sudah di kirim ke server
	// // kemudian kita taruh di server kita
	// // io package
	// io.Copy(outFile, file)

	err = ctx.SaveUploadedFile(header, fileLocation)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	// Kita akan mengekstrak dari key "user"
	// Ekstrak mnrnggunakan json.Unmarshal
	// Ekstrak ke userCredential
	var userCredential model.UserCredential

	// Masukkan datanya ke struct userCredential
	json.Unmarshal([]byte(user), &userCredential)

	// simpan ke slice
	// ambil nama filenya
	userCredential.Photo = fileLocation
	UCDb = append(UCDb, userCredential)

	model.SendSingleResponse(ctx, "berhasil update", userCredential)
}
