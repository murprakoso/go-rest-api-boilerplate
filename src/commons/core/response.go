package core

import "github.com/gin-gonic/gin"

// ResponseError sends an error response with the specified status code and message.
func ResponseError(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, gin.H{"error": message})
}

// ResponseJSON sends a JSON response with the specified status code and data.
func ResponseJSON(c *gin.Context, statusCode int, success bool, message string, data interface{}) {
	c.JSON(statusCode, gin.H{
		"success": success,
		"message": message,
		"data":    data,
	})
}
