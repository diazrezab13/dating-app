package main

import (
	"dating-app/config"
	"dating-app/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Uncomment code below to Initialize the database and Redis connection
	config.ConnectDatabase()
	config.ConnectRedis()

	// Setup routes
	routes.SetupRoutes(r)

	// Start the server on port 8080
	r.Run(":8080")
}
