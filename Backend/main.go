package main

import (
	"github.com/Abdelhamid-Mouloudi/go_employee_system_backend/config"
	"github.com/Abdelhamid-Mouloudi/go_employee_system_backend/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()

	r := gin.Default()
	routes.EmployeeRoutes(r)

	r.Run(":8080") // Start server on port 8080
}
