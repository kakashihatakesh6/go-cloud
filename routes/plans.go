package routes

import (
	"go-cloud/handlers"

	"github.com/gin-gonic/gin"
)

func SetupPlans(r *gin.RouterGroup){
	plan := r.Group("/plans")

	{
		plan.POST("/create", handlers.CreatePlan)
	}
}