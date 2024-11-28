package main

import (
	"log"

	"github.com/Abdelhamid-Mouloudi/go_employee_system_backend/config"
	"github.com/Abdelhamid-Mouloudi/go_employee_system_backend/controllers"
	"github.com/Abdelhamid-Mouloudi/go_employee_system_backend/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// Connect to MongoDB
	config.ConnectDB()
	log.Println("Database connected")

	// Initialize employee controller
	controllers.InitEmployeeController()
	log.Println("Controllers initialized")

	// Set up routes
	r := gin.Default()
	routes.EmployeeRoutes(r)
	log.Println("Routes configured")

	// Start server on port 8080
	log.Println("Server running on port 8080")
	r.Run(":8080")
}
