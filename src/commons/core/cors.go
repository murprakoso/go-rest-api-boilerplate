package core

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// SetupCors mengatur middleware CORS untuk aplikasi Gin.
func SetupCors() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	config.AddAllowHeaders("Authorization")
	config.AddExposeHeaders("Content-Length")
	return cors.New(config)
}
