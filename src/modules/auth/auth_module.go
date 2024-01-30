package auth

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var authHandler *SAuthHandler

func InitRouterGroup(router *gin.RouterGroup, db *gorm.DB) {
	// Init Module
	authRepository := NewAuthRepository(db)
	authService := NewAuthService(authRepository)
	authHandler = NewAuthHandler(authService)

	// Init Router
	router.POST("/register", authHandler.Register)
	router.POST("/login", authHandler.Login)
}
