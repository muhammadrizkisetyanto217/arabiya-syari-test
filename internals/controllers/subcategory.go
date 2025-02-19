package controllers

import (
	"log"
	"net/http"
	"strconv"

	"arabiya-syari/internals/database"
	"arabiya-syari/internals/models"

	"github.com/gin-gonic/gin"
)

type SubcategoryController struct{}



//* CREATE SUBCATEGORIES
func (sc *SubcategoryController) CreateSubcategory(c *gin.Context) {
	var subcategory models.Subcategory

	// Log request payload
	log.Println("[INFO] Received request to create subcategory")

	// Bind JSON dan cek error
	if err := c.ShouldBindJSON(&subcategory); err != nil {
		log.Printf("[ERROR] Failed to parse JSON: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}
	log.Printf("[DEBUG] Parsed subcategory: %+v\n", subcategory)

	// Simpan ke database
	if err := database.DB.Create(&subcategory).Error; err != nil {
		log.Printf("[ERROR] Failed to insert subcategory: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create subcategory"})
		return
	}

	// Log keberhasilan insert
	log.Printf("[SUCCESS] Subcategory created successfully: ID=%d, Name=%s\n", subcategory.ID, subcategory.Name)
	c.JSON(http.StatusCreated, gin.H{"message": "Subcategory created", "data": subcategory})
}



//* GET ALL CATEGORIES
func (sc *SubcategoryController) GetSubcategories(c *gin.Context) {
	var subcategories []models.Subcategory

	log.Println("[INFO] Received request to fetch all subcategories")

	// Query database
	if err := database.DB.Find(&subcategories).Error; err != nil {
		log.Printf("[ERROR] Failed to fetch subcategories: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve subcategories"})
		return
	}

	// Log jumlah data
	log.Printf("[SUCCESS] Retrieved %d subcategories\n", len(subcategories))
	c.JSON(http.StatusOK, gin.H{"data": subcategories})
}

//* GET SINGLE SUBCATEGORY
func (sc *SubcategoryController) GetSubcategory(c *gin.Context) {
	var subcategory models.Subcategory
	id := c.Param("id")

	log.Printf("[INFO] Fetching subcategory with ID: %s\n", id)

	// Validasi ID harus angka
	if _, err := strconv.Atoi(id); err != nil {
		log.Printf("[ERROR] Invalid subcategory ID: %s\n", id)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid subcategory ID"})
		return
	}

	// Query database
	if err := database.DB.First(&subcategory, id).Error; err != nil {
		log.Printf("[ERROR] Subcategory with ID %s not found\n", id)
		c.JSON(http.StatusNotFound, gin.H{"error": "Subcategory not found"})
		return
	}

	// Log sukses
	log.Printf("[SUCCESS] Retrieved subcategory: ID=%d, Name=%s\n", subcategory.ID, subcategory.Name)
	c.JSON(http.StatusOK, gin.H{"data": subcategory})
}

//* UPDATE SUBCATEGORY
func (sc *SubcategoryController) UpdateSubcategory(c *gin.Context) {
	var subcategory models.Subcategory
	id := c.Param("id")

	// Cek apakah data ada
	if err := database.DB.First(&subcategory, id).Error; err != nil {
		log.Printf("Subcategory with ID %s not found: %v", id, err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Subcategory not found"})
		return
	}

	// Bind JSON ke struct baru agar hanya mengupdate field tertentu
	var input models.Subcategory
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Printf("Invalid JSON input: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update hanya field yang dikirim
	if err := database.DB.Model(&subcategory).Updates(input).Error; err != nil {
		log.Printf("Failed to update subcategory: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update subcategory"})
		return
	}

	log.Printf("Subcategory with ID %s updated successfully", id)
	c.JSON(http.StatusOK, gin.H{"message": "Subcategory updated", "data": subcategory})
}


//* DELETE SUBCATEGORY
func (sc *SubcategoryController) DeleteSubcategory(c *gin.Context) {
	var subcategory models.Subcategory
	id := c.Param("id")

	// Cari data berdasarkan ID
	if err := database.DB.First(&subcategory, id).Error; err != nil {
		log.Printf("Subcategory with ID %s not found: %v", id, err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Subcategory not found"})
		return
	}

	// Gunakan transaksi agar lebih aman
	tx := database.DB.Begin()
	if err := tx.Delete(&subcategory).Error; err != nil {
		tx.Rollback()
		log.Printf("Failed to delete subcategory: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete subcategory"})
		return
	}
	tx.Commit()

	log.Printf("Subcategory with ID %s deleted successfully", id)
	c.JSON(http.StatusOK, gin.H{"message": "Subcategory deleted"})
}


//* GET SUBCATEGORIES BY CATEGORY
func (sc *SubcategoryController) GetSubcategoriesByCategory(c *gin.Context) {
	categoryID := c.Param("id")
	var subcategories []models.Subcategory

	// Pastikan categoryID valid
	if categoryID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Category ID is required"})
		return
	}

	// Query data
	if err := database.DB.Where("category_id = ?", categoryID).Find(&subcategories).Error; err != nil {
		log.Printf("Error fetching subcategories for category %s: %v", categoryID, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch subcategories"})
		return
	}

	// Jika kosong, kembalikan array kosong daripada null
	if len(subcategories) == 0 {
		c.JSON(http.StatusOK, gin.H{"data": []models.Subcategory{}})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": subcategories})
}
