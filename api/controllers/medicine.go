package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/faridlamaul/medlit-api-backend/api/models"
)

type MedicineInput struct {
	GenericName 	 string `json:"generic_name" binding:"required"`
	PhotoURL  		 string `json:"photo_url" binding:"required"`
	Purpose 		 string `json:"purpose" binding:"required"`
	SideEffects 	 string `json:"side_effects" binding:"required"`
	Contraindication string `json:"contraindication" binding:"required"`
}

func AddMedicine(c *gin.Context) {
	var input MedicineInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "true",
			"message": err.Error(),
		})
		return
	}

	medicine := models.Medicine{}
	medicine.GenericName = input.GenericName
	medicine.PhotoURL = input.PhotoURL
	medicine.Purpose = input.Purpose
	medicine.SideEffects = input.SideEffects
	medicine.Contraindication = input.Contraindication

	_, err := medicine.CreateMedicine()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "true",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error": "false",
		"message": "Medicine created",
	})
}

func GetAllMedicine(c *gin.Context) {
	medicine := models.GetAllMedicine()

	c.JSON(http.StatusOK, gin.H{
		"error": "false",
		"message": "Medicine list",
		"medicineList": medicine,
	})
}

func GetMedicineByID(c *gin.Context) {
	id := c.Param("id")

	medicine, err := models.GetMedicineByID(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "true",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error": "false",
		"message": "Medicine found",
		"medicine": medicine,
	})
}

func GetMedicineByQuery(c *gin.Context) {
	generic_name := c.Query("generic_name")

	medicine, err := models.GetMedicineByQuery(generic_name)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "true",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error": "false",
		"message": "Medicine found",
		"medicineList": medicine,
	})
}