package main

import (
	"github.com/gin-gonic/gin"
	"github.com/your_project_path/config"
	"github.com/your_project_path/routes"
)

func main() {
	config.ConnectDB()

	r := gin.Default()
	routes.EmployeeRoutes(r)

	r.Run(":8080") // Start server on port 8080
}
