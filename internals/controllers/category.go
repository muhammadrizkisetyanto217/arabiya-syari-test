package controllers

import (
	"net/http"

	"arabiya-syari/internals/database"
	"arabiya-syari/internals/models"

	"github.com/gin-gonic/gin"
)

type CategoryController struct{}

// Create Category
func (cc *CategoryController) CreateCategory(c *gin.Context) {
	var category models.Category

	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Create(&category)
	c.JSON(http.StatusCreated, gin.H{"message": "Category created", "data": category})
}

// Get All Categories
func (cc *CategoryController) GetCategories(c *gin.Context) {
	var categories []models.Category
	database.DB.Find(&categories)

	c.JSON(http.StatusOK, gin.H{"data": categories})
}

// Get Single Category by ID
func (cc *CategoryController) GetCategory(c *gin.Context) {
	var category models.Category
	id := c.Param("id")

	if err := database.DB.First(&category, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": category})
}

// Update Category
func (cc *CategoryController) UpdateCategory(c *gin.Context) {
	var category models.Category
	id := c.Param("id")

	if err := database.DB.First(&category, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}

	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Save(&category)
	c.JSON(http.StatusOK, gin.H{"message": "Category updated", "data": category})
}

// Delete Category
func (cc *CategoryController) DeleteCategory(c *gin.Context) {
	var category models.Category
	id := c.Param("id")

	if err := database.DB.First(&category, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}

	database.DB.Delete(&category)
	c.JSON(http.StatusOK, gin.H{"message": "Category deleted"})
}

// Get Categories by Difficulty
func (cc *CategoryController) GetCategoriesByDifficulty(c *gin.Context) {
	difficultyID := c.Param("id")
	var categories []models.Category

	if err := database.DB.Where("difficulty_id = ?", difficultyID).Find(&categories).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": categories})
}
