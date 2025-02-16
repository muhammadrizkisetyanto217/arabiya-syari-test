package main

import (
	"arabiya-syari/internals/handlers/memberships"
	// "net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
    //* Sama seperti 
    // r := gin.New()
    // r.Use(gin.Logger())
    // r.Use(gin.Recovery())
    //? 1. r.Use(gin.Logger()) → Middleware Logging
    // Middleware ini mencatat setiap permintaan HTTP yang masuk, termasuk metode (GET, POST, dll.), URL, status kode, dan waktu eksekusi.
    // Berguna untuk debugging dan memantau lalu lintas aplikasi.
    //? 2. r.Use(gin.Recovery()) → Middleware Recovery
    // Middleware ini menangani panic yang terjadi selama eksekusi request.
    // Jika ada error fatal (panic), middleware ini akan menangkapnya dan mencegah server crash.
    // Sebagai gantinya, server akan mengembalikan respons HTTP 500.


	//? Inisialisasi handler memberships
	membershipHandler := memberships.NewHandler(r)

	//? Registrasi route
	membershipHandler.RegisterRoutes()

    //* Cara sekali pakai
    // memberships.NewHandler(r).RegisterRoutes();

	// Jalankan server
	r.Run(":8080") // Server berjalan di port 8080
}
