package controllers

import (
	"net/http"

	"arabiya-syari/internals/database"
	"arabiya-syari/internals/models"

	"github.com/gin-gonic/gin"
)

type SubcategoryController struct{}

// Create Subcategory
func (sc *SubcategoryController) CreateSubcategory(c *gin.Context) {
	var subcategory models.Subcategory

	if err := c.ShouldBindJSON(&subcategory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Create(&subcategory)
	c.JSON(http.StatusCreated, gin.H{"message": "Subcategory created", "data": subcategory})
}

// Get All Subcategories
func (sc *SubcategoryController) GetSubcategories(c *gin.Context) {
	var subcategories []models.Subcategory
	database.DB.Find(&subcategories)

	c.JSON(http.StatusOK, gin.H{"data": subcategories})
}

// Get Single Subcategory by ID
func (sc *SubcategoryController) GetSubcategory(c *gin.Context) {
	var subcategory models.Subcategory
	id := c.Param("id")

	if err := database.DB.First(&subcategory, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Subcategory not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": subcategory})
}

// Update Subcategory
func (sc *SubcategoryController) UpdateSubcategory(c *gin.Context) {
	var subcategory models.Subcategory
	id := c.Param("id")

	if err := database.DB.First(&subcategory, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Subcategory not found"})
		return
	}

	if err := c.ShouldBindJSON(&subcategory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Save(&subcategory)
	c.JSON(http.StatusOK, gin.H{"message": "Subcategory updated", "data": subcategory})
}

// Delete Subcategory
func (sc *SubcategoryController) DeleteSubcategory(c *gin.Context) {
	var subcategory models.Subcategory
	id := c.Param("id")

	if err := database.DB.First(&subcategory, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Subcategory not found"})
		return
	}

	database.DB.Delete(&subcategory)
	c.JSON(http.StatusOK, gin.H{"message": "Subcategory deleted"})
}

// Get Subcategories by Category ID
func (sc *SubcategoryController) GetSubcategoriesByCategory(c *gin.Context) {
	categoryID := c.Param("id")
	var subcategories []models.Subcategory

	if err := database.DB.Where("category_id = ?", categoryID).Find(&subcategories).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": subcategories})
}
