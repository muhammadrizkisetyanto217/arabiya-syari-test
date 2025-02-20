package routes

import (
	"arabiya-syari/internals/controllers"
	// "arabiya-syari/internals/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)


}
