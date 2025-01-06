package main

import (
	"log"
	"os"

	"go-cloud/config"
	"go-cloud/database"
	"go-cloud/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)


func main() {
	envFile := ".env"
	if _, err := os.Stat(".env.local"); err == nil {
		envFile = ".env.local"
	}

	err := godotenv.Load(envFile)
	if err != nil {
		log.Fatalf("Error loading %s file: %s", envFile, err)
	}

	if os.Getenv("ENV") == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	config.LoadConfig()

	database.ConnectToPostgres(
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	r := gin.Default()

	r.Use(func(c *gin.Context)  {
		c.Set("db", database.DB)
		c.Next()
	})

	api := r.Group("/api")

	routes.SetupAuthRoutes(api)
	routes.SetupPlans(api)
	routes.SetupComputeResource(api)
	
	log.Println("Starting server on :8181")
	if err := r.Run(":8181"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}


}
