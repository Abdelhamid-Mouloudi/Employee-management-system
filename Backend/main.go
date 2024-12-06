package main

import (
	"log"

	"github.com/Abdelhamid-Mouloudi/go_employee_system_backend/config"
	"github.com/Abdelhamid-Mouloudi/go_employee_system_backend/controllers"
	"github.com/Abdelhamid-Mouloudi/go_employee_system_backend/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Connect to MongoDB
	config.ConnectDB()
	log.Println("Database connected")

	// Initialize employee controller
	controllers.InitEmployeeController()
	log.Println("Controllers initialized")

	// Set up Gin router
	r := gin.Default()

	// Add CORS middleware
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},                                       // Frontend URL
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, // Allowed HTTP methods
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"}, // Allowed headers
		ExposeHeaders:    []string{"Content-Length"},                          // Expose additional headers to the client
		AllowCredentials: true,                                                // Allow cookies and authentication headers
	}))

	// Configure routes
	routes.EmployeeRoutes(r)
	log.Println("Routes configured")

	// Start server on port 8080
	log.Println("Server running on port 8080")
	r.Run(":8080")
}
