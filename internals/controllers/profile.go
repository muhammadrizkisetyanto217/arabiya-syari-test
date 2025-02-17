package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"arabiya-syari/internals/models"
)

func Profile(c *gin.Context) {
	// Ambil data user dari middleware
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Konversi user ke struct yang sesuai
	u, ok := user.(models.User)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User data corrupted"})
		return
	}

	// Kirim respons ke frontend
	c.JSON(http.StatusOK, gin.H{
		"id":       u.ID,
		"email":    u.Email,
		"created":  u.CreatedAt,
	})
}
