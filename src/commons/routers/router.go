package routers

import (
	"github.com/gin-gonic/gin"
	"go-rest-api-boilerplate/src/commons/core"
	"go-rest-api-boilerplate/src/modules/product"
)

func Router() *gin.Engine {
	router := gin.Default()

	//API Version
	v1 := router.Group("/v1")
	//Route
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Selamat datang!",
		})
	})

	v1.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Selamat datang! V1 API",
		})
	})

	v1.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	//ProductHandler
	productRepository := product.NewProductRepository(core.DB)
	productService := product.NewProductService(productRepository)
	productHandler := product.NewProductHandler(productService)

	v1.GET("/product", productHandler.ShowProducts)
	v1.GET("/product/:id", productHandler.ShowProduct)
	v1.POST("/product", productHandler.CreateProduct)

	return router
}
