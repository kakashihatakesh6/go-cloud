package routes

import (
	"go-cloud/handlers"

	"github.com/gin-gonic/gin"
)

func SetupAuthRoutes(r *gin.RouterGroup) {
	auth := r.Group("/auth")
	
	{
		auth.POST("/signup", handlers.Signup)
		auth.POST("/login", handlers.Login)
	}
}