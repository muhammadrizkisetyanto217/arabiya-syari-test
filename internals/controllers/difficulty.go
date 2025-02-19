package controllers

import (
	"log"
	"net/http"
	"strconv"

	"arabiya-syari/internals/database"
	"arabiya-syari/internals/models"

	"github.com/gin-gonic/gin"
)

//* CREATE DIFFICULTY
//? c *gin.Context adalah parameter yang mewakili konteks permintaan HTTP (digunakan untuk membaca input dan mengirim respons).
func CreateDifficulty(c *gin.Context) {
	var difficulty models.Difficulty

	log.Println("[INFO] Received request to create difficulty")

	// Parsing JSON
	if err := c.ShouldBindJSON(&difficulty); err != nil {
		log.Printf("[ERROR] Failed to parse JSON: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}
	log.Printf("[DEBUG] Parsed difficulty: %+v\n", difficulty)

	// Insert ke database
	if err := database.DB.Create(&difficulty).Error; err != nil {
		log.Printf("[ERROR] Failed to insert difficulty: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create difficulty"})
		return
	}

	// Log sukses
	log.Printf("[SUCCESS] Difficulty created: ID=%d, Level=%s\n", difficulty.ID, difficulty.Name)
	c.JSON(http.StatusCreated, gin.H{"message": "Difficulty created", "data": difficulty})
}

//* GET ALL DIFFICULTIES
func GetDifficulties(c *gin.Context) {
	var difficulties []models.Difficulty

	log.Println("[INFO] Received request to fetch all difficulties")

	// Query database
	if err := database.DB.Find(&difficulties).Error; err != nil {
		log.Printf("[ERROR] Failed to fetch difficulties: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve difficulties"})
		return
	}

	// Log jumlah data
	log.Printf("[SUCCESS] Retrieved %d difficulties\n", len(difficulties))
	c.JSON(http.StatusOK, gin.H{"data": difficulties})
}


//* GET SINGLE DIFFICULTY
func GetDifficultyByID(c *gin.Context) {
	var difficulty models.Difficulty
	id := c.Param("id")

	log.Printf("[INFO] Fetching difficulty with ID: %s\n", id)

	// Validasi ID harus angka
	if _, err := strconv.Atoi(id); err != nil {
		log.Printf("[ERROR] Invalid difficulty ID: %s\n", id)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid difficulty ID"})
		return
	}

	// Query database
	if err := database.DB.First(&difficulty, id).Error; err != nil {
		log.Printf("[ERROR] Difficulty with ID %s not found\n", id)
		c.JSON(http.StatusNotFound, gin.H{"error": "Difficulty not found"})
		return
	}

	// Log sukses
	log.Printf("[SUCCESS] Retrieved difficulty: ID=%d, Name=%s\n", difficulty.ID, difficulty.Name)
	c.JSON(http.StatusOK, gin.H{"data": difficulty})
}


//* UPDATE DIFFICULTY
func UpdateDifficulty(c *gin.Context) {
	var difficulty models.Difficulty
	id := c.Param("id")

	// Cek apakah data ada
	if err := database.DB.First(&difficulty, id).Error; err != nil {
		log.Printf("Difficulty with ID %s not found: %v", id, err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Difficulty not found"})
		return
	}

	// Bind JSON ke struct baru agar hanya mengupdate field tertentu
	var input models.Difficulty
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Printf("Invalid JSON input: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update hanya field yang dikirim
	if err := database.DB.Model(&difficulty).Updates(input).Error; err != nil {
		log.Printf("Failed to update difficulty: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update difficulty"})
		return
	}

	log.Printf("Difficulty with ID %s updated successfully", id)
	c.JSON(http.StatusOK, gin.H{"message": "Difficulty updated", "data": difficulty})
}


//* DELETE DIFFICULTY
func DeleteDifficulty(c *gin.Context) {
	var difficulty models.Difficulty
	id := c.Param("id")

	// Cari data berdasarkan ID
	if err := database.DB.First(&difficulty, id).Error; err != nil {
		log.Printf("Difficulty with ID %s not found: %v", id, err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Difficulty not found"})
		return
	}

	// Hapus data
	if err := database.DB.Delete(&difficulty).Error; err != nil {
		log.Printf("Failed to delete difficulty: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete difficulty"})
		return
	}

	log.Printf("Difficulty with ID %s deleted successfully", id)
	c.JSON(http.StatusOK, gin.H{"message": "Difficulty deleted"})
}
