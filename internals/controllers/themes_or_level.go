package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"arabiya-syari/internals/database"
	"arabiya-syari/internals/models"
)

type ThemeOrLevelController struct{}

//* CREATE THEME OR LEVEL
func (tlc *ThemeOrLevelController) CreateThemeOrLevel(c *gin.Context) {
	var themeOrLevel models.ThemesOrLevel

	log.Println("[INFO] Received request to create ThemeOrLevel")

	if err := c.ShouldBindJSON(&themeOrLevel); err != nil {
		log.Printf("[ERROR] Failed to parse JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	log.Printf("[DEBUG] Parsed ThemeOrLevel: %+v", themeOrLevel)

	if err := database.DB.Create(&themeOrLevel).Error; err != nil {
		log.Printf("[ERROR] Failed to insert ThemeOrLevel: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create ThemeOrLevel"})
		return
	}

	log.Printf("[SUCCESS] ThemeOrLevel created: ID=%d, Name=%s", themeOrLevel.ID, themeOrLevel.Name)
	c.JSON(http.StatusCreated, gin.H{"message": "ThemeOrLevel created", "data": themeOrLevel})
}

//* GET ALL THEMES OR LEVELS
func (tlc *ThemeOrLevelController) GetThemesOrLevels(c *gin.Context) {
	var themesOrLevels []models.ThemesOrLevel

	log.Println("[INFO] Received request to fetch all   bThemesOrLevels")

	if err := database.DB.Find(&themesOrLevels).Error; err != nil {
		log.Printf("[ERROR] Failed to fetch ThemesOrLevels: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve ThemesOrLevels"})
		return
	}

	log.Printf("[SUCCESS] Retrieved %d ThemesOrLevels", len(themesOrLevels))
	c.JSON(http.StatusOK, gin.H{"data": themesOrLevels})
}

//* GET SINGLE THEME OR LEVEL
func (tlc *ThemeOrLevelController) GetThemeOrLevel(c *gin.Context) {
	var themeOrLevel models.ThemesOrLevel
	id := c.Param("id")

	log.Printf("[INFO] Received request to fetch ThemeOrLevel with ID: %s", id)

	if _, err := strconv.Atoi(id); err != nil {
		log.Printf("[ERROR] Invalid ThemeOrLevel ID: %s", id)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ThemeOrLevel ID"})
		return
	}

	if err := database.DB.First(&themeOrLevel, id).Error; err != nil {
		log.Printf("[ERROR] ThemeOrLevel with ID %s not found", id)
		c.JSON(http.StatusNotFound, gin.H{"error": "ThemeOrLevel not found"})
		return
	}

	log.Printf("[SUCCESS] Retrieved ThemeOrLevel: ID=%d, Name=%s", themeOrLevel.ID, themeOrLevel.Name)
	c.JSON(http.StatusOK, gin.H{"data": themeOrLevel})
}

//* UPDATE THEME OR LEVEL
func (tlc *ThemeOrLevelController) UpdateThemeOrLevel(c *gin.Context) {
	var themeOrLevel models.ThemesOrLevel
	id := c.Param("id")

	log.Printf("[INFO] Received request to update ThemeOrLevel with ID: %s", id)

	if err := database.DB.First(&themeOrLevel, id).Error; err != nil {
		log.Printf("[ERROR] ThemeOrLevel with ID %s not found", id)
		c.JSON(http.StatusNotFound, gin.H{"error": "ThemeOrLevel not found"})
		return
	}

	var input models.ThemesOrLevel
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Printf("[ERROR] Invalid JSON input: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Model(&themeOrLevel).Updates(input).Error; err != nil {
		log.Printf("[ERROR] Failed to update ThemeOrLevel: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update ThemeOrLevel"})
		return
	}

	log.Printf("[SUCCESS] ThemeOrLevel with ID %s updated successfully", id)
	c.JSON(http.StatusOK, gin.H{"message": "ThemeOrLevel updated", "data": themeOrLevel})
}

//* DELETE THEME OR LEVEL
func (tlc *ThemeOrLevelController) DeleteThemeOrLevel(c *gin.Context) {
	var themeOrLevel models.ThemesOrLevel
	id := c.Param("id")

	log.Printf("[INFO] Received request to delete ThemeOrLevel with ID: %s", id)

	if err := database.DB.First(&themeOrLevel, id).Error; err != nil {
		log.Printf("[ERROR] ThemeOrLevel with ID %s not found", id)
		c.JSON(http.StatusNotFound, gin.H{"error": "ThemeOrLevel not found"})
		return
	}

	tx := database.DB.Begin()
	if err := tx.Delete(&themeOrLevel).Error; err != nil {
		tx.Rollback()
		log.Printf("[ERROR] Failed to delete ThemeOrLevel: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete ThemeOrLevel"})
		return
	}
	tx.Commit()

	log.Printf("[SUCCESS] ThemeOrLevel with ID %s deleted successfully", id)
	c.JSON(http.StatusOK, gin.H{"message": "ThemeOrLevel deleted"})
}
