package database

import (
	"log"
	"os"

	// "github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// func ConnectDB() (*gorm.DB, error) {
// 	// Load environment variables
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Println("No .env file found")
// 	}

// 	dsn := os.Getenv("DATABASE_URL")
// 	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		return nil, err
// 	}

// 	log.Println("Connected to PostgreSQL successfully")
// 	return db, nil
// }

var DB *gorm.DB

func ConnectDatabase() {
	dsn := os.Getenv("DATABASE_URL") // Simpan koneksi di ENV
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	DB = db
	log.Println("Database connected")
}