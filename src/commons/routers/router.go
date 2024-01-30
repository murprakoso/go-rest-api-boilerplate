package routers

import (
	"github.com/gin-gonic/gin"
	"go-rest-api-boilerplate/src/commons/core"
	"go-rest-api-boilerplate/src/commons/middleware"
	"go-rest-api-boilerplate/src/modules/auth"
	"go-rest-api-boilerplate/src/modules/product"
	"go-rest-api-boilerplate/src/modules/unit"
)

func Router() *gin.Engine {
	// Create a new Gin router with default middleware
	router := gin.Default()

	// API Version Grouping
	v1 := router.Group("/v1")

	// Mendaftarkan middleware ke grup router untuk validasi ID parameter
	v1.Use(middleware.ValidateIDParamMiddleware("id"))

	// Default route
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Welcome to the API!"})
	})

	// API Version-specific route
	v1.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Welcome to the V1 API!"})
	})

	// List of imported and configured router groups (add new router groups here)
	//
	// Imported Router Groups:
	// - product.SetProductRouterGroup(v1)

	// Import and configure the product routes within the v1 group
	product.InitRouterGroup(v1, core.DB)
	unit.InitRouterGroup(v1, core.DB)
	auth.InitRouterGroup(v1, core.DB)

	return router
}
