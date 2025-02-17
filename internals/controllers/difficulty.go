package controllers

import (
	"net/http"

	"arabiya-syari/internals/database"
	"arabiya-syari/internals/models"

	"github.com/gin-gonic/gin"
)

// Create Difficulty
func CreateDifficulty(c *gin.Context) {
	var difficulty models.Difficulty

	if err := c.ShouldBindJSON(&difficulty); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Create(&difficulty)
	c.JSON(http.StatusCreated, gin.H{"message": "Difficulty created", "data": difficulty})
}

// Get All Difficulties
func GetDifficulties(c *gin.Context) {
	var difficulties []models.Difficulty
	database.DB.Find(&difficulties)

	c.JSON(http.StatusOK, gin.H{"data": difficulties})
}

// Get Single Difficulty by ID
func GetDifficultyByID(c *gin.Context) {
	var difficulty models.Difficulty
	id := c.Param("id")

	if err := database.DB.First(&difficulty, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Difficulty not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": difficulty})
}

// Update Difficulty
func UpdateDifficulty(c *gin.Context) {
	var difficulty models.Difficulty
	id := c.Param("id")

	if err := database.DB.First(&difficulty, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Difficulty not found"})
		return
	}

	if err := c.ShouldBindJSON(&difficulty); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Save(&difficulty)
	c.JSON(http.StatusOK, gin.H{"message": "Difficulty updated", "data": difficulty})
}

// Delete Difficulty
func DeleteDifficulty(c *gin.Context) {
	var difficulty models.Difficulty
	id := c.Param("id")

	if err := database.DB.First(&difficulty, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Difficulty not found"})
		return
	}

	database.DB.Delete(&difficulty)
	c.JSON(http.StatusOK, gin.H{"message": "Difficulty deleted"})
}
