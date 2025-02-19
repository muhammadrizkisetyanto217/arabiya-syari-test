package main

import (
	"arabiya-syari/internals/database"
	"arabiya-syari/internals/routes"
	"os"

	// "arabiya-syari/internals/handlers/memberships"

	// "fmt"
	"log"

	// "net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	// "gorm.io/gorm"
)


func main() {

    // Koneksi ke database
   // Load .env
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	database.ConnectDatabase()



	// r := gin.Default()
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


    // r := routes.SetupRouter()
	// Buat instance Gin dengan `gin.Default()`
	r := gin.Default()

    r.Use(func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

        // Handle preflight request
        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }

        c.Next()
    })

	// Register routes
	routes.SetupRouter(r)

        r.Use(func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

        // Handle preflight request
        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }

        c.Next()
    })

    // Category routes
    routes.CategoryRouter(r)

	// Set port default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Run server
	r.Run(":" + port)


	//? Inisialisasi handler memberships
	// membershipHandler := memberships.NewHandler(r)

	// //? Registrasi route
	// membershipHandler.RegisterRoutes()

    //* Cara sekali pakai
    // memberships.NewHandler(r).RegisterRoutes();

	// Jalankan server
	r.Run(":8080") // Server berjalan di port 8080
}
