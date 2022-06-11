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

	api := router.Group("/api")

	api.GET("/", func (c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to Medlit API",
		})
	})
	api.POST("/medlit/login", controllers.Login)
	api.POST("/medlit/register", controllers.Register)
	api.POST("/medlit/medicine/add", controllers.AddMedicine)
	api.GET("/medlit/medicine/get/all", controllers.GetAllMedicine)
	api.GET("/medlit/medicine/get/:id", controllers.GetMedicineByID)
	api.GET("/medlit/medicine/search", controllers.GetMedicineByQuery)

	router.Run(":8080")
}