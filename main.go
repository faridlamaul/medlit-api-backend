package main

import (
	"net/http"

	"github.com/faridlamaul/medlit-api-backend/api/controllers"
	"github.com/faridlamaul/medlit-api-backend/api/models"
	"github.com/gin-gonic/gin"
)

func main() {
	models.ConnectToDatabase()

	router := gin.Default()

	api := router.Group("/api/medlit")

	api.GET("/", func (c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"error": "false",
			"message": "Welcome to Medlit API",
		})
	})
	api.POST("/login", controllers.Login)
	api.POST("/register", controllers.Register)
	api.POST("/medicine/add", controllers.AddMedicine)
	api.GET("/medicine/get/all", controllers.GetAllMedicine)
	api.GET("/medicine/get/:id", controllers.GetMedicineByID)
	api.GET("/medicine/search", controllers.GetMedicineByQuery)

	router.Run(":8080")
}