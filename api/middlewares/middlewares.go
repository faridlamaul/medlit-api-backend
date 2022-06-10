package middlewares

import (
	"github.com/faridlamaul/medlit-api-backend/api/utils/token"
	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := token.TokenValid(c)
		if err != nil {
			c.JSON(401, gin.H{
				"message": err.Error(),
			})
			c.Abort()
			return
		}
		c.Next()
	}
}