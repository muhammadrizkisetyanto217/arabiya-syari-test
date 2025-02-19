package controllers

import (
	"log"
	"net/http"
	"strconv"

	"arabiya-syari/internals/database"
	"arabiya-syari/internals/models"

	"github.com/gin-gonic/gin"
)

type CategoryController struct{}

//* CREATE CATEGORY
func (cc *CategoryController) CreateCategory(c *gin.Context) {
	var category models.Category

	log.Println("[INFO] Received request to create category")

	// Parsing JSON
	if err := c.ShouldBindJSON(&category); err != nil {
		log.Printf("[ERROR] Failed to parse JSON: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}
	log.Printf("[DEBUG] Parsed category: %+v\n", category)

	// Insert ke database
	if err := database.DB.Create(&category).Error; err != nil {
		log.Printf("[ERROR] Failed to insert category: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create category"})
		return
	}

	// Log sukses
	log.Printf("[SUCCESS] Category created: ID=%d, Name=%s\n", category.ID, category.Name)
	c.JSON(http.StatusCreated, gin.H{"message": "Category created", "data": category})
}

//* GET ALL CATEGORIES
func (cc *CategoryController) GetCategories(c *gin.Context) {
	var categories []models.Category

	log.Println("[INFO] Received request to fetch all categories")

	// Query database
	if err := database.DB.Find(&categories).Error; err != nil {
		log.Printf("[ERROR] Failed to fetch categories: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve categories"})
		return
	}

	// Log jumlah data
	log.Printf("[SUCCESS] Retrieved %d categories\n", len(categories))
	c.JSON(http.StatusOK, gin.H{"data": categories})
}


//* GET SINGLE CATEGORY
func (cc *CategoryController) GetCategory(c *gin.Context) {
	var category models.Category
	id := c.Param("id")

	log.Printf("[INFO] Received request to fetch category with ID: %s\n", id)

	// Validasi apakah ID dalam bentuk angka
	if _, err := strconv.Atoi(id); err != nil {
		log.Printf("[ERROR] Invalid category ID: %s\n", id)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
		return
	}

	// Query database
	if err := database.DB.First(&category, id).Error; err != nil {
		log.Printf("[ERROR] Category with ID %s not found\n", id)
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}

	// Log sukses
	log.Printf("[SUCCESS] Retrieved category: ID=%d, Name=%s\n", category.ID, category.Name)
	c.JSON(http.StatusOK, gin.H{"data": category})
}


//* UPDATE CATEGORY
func (cc *CategoryController) UpdateCategory(c *gin.Context) {
	var category models.Category
	id := c.Param("id")

	// Cek apakah data ada
	if err := database.DB.First(&category, id).Error; err != nil {
		log.Printf("Category with ID %s not found: %v", id, err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}

	// Bind JSON ke struct baru agar hanya mengupdate field tertentu
	var input models.Category
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Printf("Invalid JSON input: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update hanya field yang dikirim
	if err := database.DB.Model(&category).Updates(input).Error; err != nil {
		log.Printf("Failed to update category: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update category"})
		return
	}

	log.Printf("Category with ID %s updated successfully", id)
	c.JSON(http.StatusOK, gin.H{"message": "Category updated", "data": category})
}




//* DELETE CATEGORY
func (cc *CategoryController) DeleteCategory(c *gin.Context) {
	var category models.Category
	id := c.Param("id")

	// Cari data berdasarkan ID
	if err := database.DB.First(&category, id).Error; err != nil {
		log.Printf("Category with ID %s not found: %v", id, err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}

	// Gunakan transaksi agar lebih aman
	tx := database.DB.Begin()
	if err := tx.Delete(&category).Error; err != nil {
		tx.Rollback()
		log.Printf("Failed to delete category: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete category"})
		return
	}
	tx.Commit()

	log.Printf("Category with ID %s deleted successfully", id)
	c.JSON(http.StatusOK, gin.H{"message": "Category deleted"})
}


//* GET CATEGORIES BY DIFFICULTY
func (cc *CategoryController) GetCategoriesByDifficulty(c *gin.Context) {
	difficultyID := c.Param("id")
	var categories []models.Category

	// Pastikan difficultyID valid
	if difficultyID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Difficulty ID is required"})
		return
	}

	// Query data
	if err := database.DB.Where("difficulty_id = ?", difficultyID).Find(&categories).Error; err != nil {
		log.Printf("Error fetching categories for difficulty %s: %v", difficultyID, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch categories"})
		return
	}

	// Jika kosong, kembalikan array kosong daripada null
	if len(categories) == 0 {
		c.JSON(http.StatusOK, gin.H{"data": []models.Category{}})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": categories})
}
