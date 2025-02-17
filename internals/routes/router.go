package routes

import (
	"github.com/gin-gonic/gin"

	"arabiya-syari/internals/controllers"
	// "arabiya-syari/internals/middlewares"
)

func SetupRouter(r *gin.Engine) {
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)


		// Routes
	api := r.Group("/api/difficulties")
	{
		api.POST("/", controllers.CreateDifficulty)
		api.GET("/", controllers.GetDifficulties)
		api.GET("/:id", controllers.GetDifficultyByID)
		api.PUT("/:id", controllers.UpdateDifficulty)
		api.DELETE("/:id", controllers.DeleteDifficulty)
	}
}
