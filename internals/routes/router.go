package routes

import (
	"arabiya-syari/internals/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	// Instance Controller
	categoryController := controllers.CategoryController{}

	// Routes for Difficulties
	apiDifficulties := r.Group("/api/difficulties")
	{
		apiDifficulties.POST("/", controllers.CreateDifficulty)
		apiDifficulties.GET("/", controllers.GetDifficulties)
		apiDifficulties.GET("/:id", controllers.GetDifficultyByID)
		apiDifficulties.PUT("/:id", controllers.UpdateDifficulty)
		apiDifficulties.DELETE("/:id", controllers.DeleteDifficulty)
	}

	// Routes for Categories
	apiCategories := r.Group("/api/categories")
	{
		apiCategories.POST("/", categoryController.CreateCategory)
		apiCategories.GET("/", categoryController.GetCategories)
		apiCategories.GET("/:id", categoryController.GetCategory)
		apiCategories.PUT("/:id", categoryController.UpdateCategory)
		apiCategories.DELETE("/:id", categoryController.DeleteCategory)
	}

	// Get Categories by Difficulty
	r.GET("/api/difficulties/:id/categories", categoryController.GetCategoriesByDifficulty)
}
