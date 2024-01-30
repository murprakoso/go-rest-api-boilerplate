package product

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var productHandler *SProductHandler

func InitRouterGroup(router *gin.RouterGroup, db *gorm.DB) {
	// Init Module
	productRepository := NewProductRepository(db)
	productService := NewProductService(productRepository)
	productHandler = NewProductHandler(productService)

	// Init Route
	router.GET("/product", productHandler.ShowProducts)
	router.GET("/product/:id", productHandler.ShowProduct)
	router.POST("/product", productHandler.CreateProduct)
	router.PUT("/product/:id", productHandler.UpdateProduct)
	router.DELETE("/product/:id", productHandler.DestroyProduct)
}
