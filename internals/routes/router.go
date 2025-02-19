package routes

import (
	"arabiya-syari/internals/controllers"
	"arabiya-syari/internals/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

		// Instance Controller
	categoryController := controllers.CategoryController{}

	// Routes for Difficulties
	apiDifficulties := r.Group("/api/difficulties")
	apiDifficulties.Use(middlewares.AuthMiddleware())
	
	
	
	{
		apiDifficulties.POST("/", controllers.CreateDifficulty)
		apiDifficulties.GET("/", controllers.GetDifficulties)
		apiDifficulties.GET("/:id", controllers.GetDifficultyByID)
		apiDifficulties.PUT("/:id", controllers.UpdateDifficulty)
		apiDifficulties.DELETE("/:id", controllers.DeleteDifficulty)
	}

	// Routes for Categories
	apiCategories := r.Group("/api/categories")
	apiCategories.Use(middlewares.AuthMiddleware())
	
	{
		apiCategories.POST("/", categoryController.CreateCategory)
		apiCategories.GET("/", categoryController.GetCategories)
		apiCategories.GET("/:id", categoryController.GetCategory)
		apiCategories.PUT("/:id", categoryController.UpdateCategory)
		apiCategories.DELETE("/:id", categoryController.DeleteCategory)
	}

	// Get Categories by Difficulty
	r.GET("/api/difficulties/:id/categories", categoryController.GetCategoriesByDifficulty)


	// Instance Controller
	subcategoryController := controllers.SubcategoryController{}


	// Subcategory routes
	apisubcategoryGroup := r.Group("/api/subcategories")
	apisubcategoryGroup.Use(middlewares.AuthMiddleware())
	{
		apisubcategoryGroup.GET("/", subcategoryController.GetSubcategories)
		apisubcategoryGroup.GET("/:id", subcategoryController.GetSubcategory)
		apisubcategoryGroup.POST("/", subcategoryController.CreateSubcategory)
		apisubcategoryGroup.PUT("/:id", subcategoryController.UpdateSubcategory)
		apisubcategoryGroup.DELETE("/:id", subcategoryController.DeleteSubcategory)
	}


	r.GET("/api/categories/:id/subcategories", subcategoryController.GetSubcategoriesByCategory)
	

}
