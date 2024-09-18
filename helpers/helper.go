package helpers

import "github.com/gin-gonic/gin"

func RespondJSON(c *gin.Context, status int, message string, data interface{}) {
	c.JSON(status, gin.H{
		"message": message,
		"data":    data,
	})
}
