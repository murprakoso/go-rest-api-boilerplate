package product

import (
	"github.com/gin-gonic/gin"
	"go-rest-api-boilerplate/src/commons/core"
)

var productHandler *ProductHandler

func InitProductModule() {
	productRepository := NewProductRepository(core.DB)
	productService := NewProductService(productRepository)
	productHandler = NewProductHandler(productService)
}

func SetProductRouterGroup(router *gin.RouterGroup) {
	router.GET("/product", productHandler.ShowProducts)
	router.GET("/product/:id", productHandler.ShowProduct)
	router.POST("/product", productHandler.CreateProduct)
	router.PUT("/product/:id", productHandler.UpdateProduct)
	router.DELETE("/product/:id", productHandler.DestroyProduct)
}
