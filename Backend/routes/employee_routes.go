package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/your_project_path/controllers"
)

func EmployeeRoutes(router *gin.Engine) {
	router.GET("/employees", controllers.GetEmployees)
	router.GET("/employees/:id", controllers.GetEmployeeByID)
	router.POST("/employees", controllers.CreateEmployee)
	router.PUT("/employees/:id", controllers.UpdateEmployee)
	router.DELETE("/employees/:id", controllers.DeleteEmployee)
}
