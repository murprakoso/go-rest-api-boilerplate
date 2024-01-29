package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func ValidateIDParamMiddleware(paramName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		paramValue := c.Param(paramName)
		if err := ValidateID(paramName, paramValue); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		c.Next()
	}
}

func ValidateID(paramName, paramValue string) error {
	if paramValue == "" {
		return fmt.Errorf("Invalid %s", paramName)
	}

	if _, err := strconv.Atoi(paramValue); err != nil {
		return fmt.Errorf("%s must be a valid integer", paramName)
	}

	return nil
}
