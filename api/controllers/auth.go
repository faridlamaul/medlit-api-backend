package controllers

import (
	"net/http"

	"github.com/faridlamaul/medlit-api-backend/api/models"
	"github.com/faridlamaul/medlit-api-backend/api/utils/token"
	"github.com/gin-gonic/gin"
)

func CurrentUser (c *gin.Context) {
	user_id, err := token.ExtractTokenID(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "true", 
			"message": err.Error(),
		})
		return
	}

	u, err := models.GetUserByID(user_id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errror": "true",
			"message": err.Error(), 
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error": "false",
		"message": "success", 
		"data": u, 
	})
}

type LoginInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var input LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "true", 
			"message": err.Error(),
		})
		return
	}

	u := models.User{}

	u.Email = input.Email
	u.Password = input.Password
	u.Name = models.GetNameByEmail(u.Email)

	token, err := models.LoginCheck(u.Email, u.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "true", 
			"message": "Invalid email or password",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error" : "false", 
		"message" : "Login Successful", 
		"loginResult" : gin.H{
			"name" : u.Name, 
			"token" : token,
		}, 
	})
}

type RegisterInput struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Register(c *gin.Context) {
	var input RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "true", 
			"message": err.Error(), 
		})
		return
	}

	u := models.User{}
	
	u.Name = input.Name
	u.Email = input.Email
	u.Password = input.Password
	
	if hashErr := u.BeforeSave(); hashErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "true", 
			"message": hashErr.Error(), 
		})
		return
	}

	_, err := u.SaveUser()
	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "true", 
			"message": err.Error(), 
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error": "false", 
		"message": "Registration Successful", 
	})
}


