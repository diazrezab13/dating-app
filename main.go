package main

import (
	"dating-app/config"
	"dating-app/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	r := gin.Default()

	if err := godotenv.Load(); err != nil {
		log.Fatal("warning: .env file not found")
	}

	// Initialize the database and Redis connection
	config.ConnectDatabase()
	config.ConnectRedis()

	// Setup routes
	routes.SetupRoutes(r)

	// Start the server on port 8080
	r.Run(":8080")
}
