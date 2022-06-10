package main

import (
	"net/http"

	"github.com/faridlamaul/medlit-api-backend/api/cloudbucket"
	"github.com/faridlamaul/medlit-api-backend/api/controllers"
	"github.com/faridlamaul/medlit-api-backend/api/models"
	"github.com/gin-gonic/gin"
)

func main() {
	models.ConnectToDatabase()

	router := gin.Default()

	v1 := router.Group("/api/v1")

	v1.GET("/", func (c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"error": "false",
			"message": "Welcome to Medlit API",
		})
	})
	v1.POST("/login", controllers.Login)
	v1.POST("/register", controllers.Register)
	v1.POST("/medicine/add", controllers.AddMedicine)
	v1.GET("/medicine/get/all", controllers.GetAllMedicine)
	v1.GET("/medicine/get/:id", controllers.GetMedicineByID)
	v1.GET("/medicine/search", controllers.GetMedicineByQuery)
	v1.POST("/images/upload", cloudbucket.HandleFileUpload)

	router.Run(":8080")
}