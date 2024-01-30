package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	commons "go-rest-api-boilerplate/src/commons/config"
	"net/http"
	"time"
)

func RequireAuth(c *gin.Context) {
	//fmt.Println("In middleware of auth")
	secretKey := commons.LoadConfig().JwtSecret
	accessToken, err := c.Cookie("Authorization")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Token tidak ada",
		})
		return
	}

	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Token tidak valid",
		})
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Token telah kedaluwarsa",
			})
			return
		}

		// Find the user with token sub
		// Temukan pengguna berdasarkan ID yang diambil dari token

		c.Set("claims", claims) // Jwt Payload
		c.Next()
	} else {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Token tidak valid",
		})
	}
}
