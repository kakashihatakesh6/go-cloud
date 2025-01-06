package routes

import (
	"go-cloud/handlers"

	"github.com/gin-gonic/gin"
)

func SetupComputeResource(r *gin.RouterGroup) {
	computeResource := r.Group("/compute-resource")

	{
		computeResource.POST("/create", handlers.CreateComputeResource)
	}
}