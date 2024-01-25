package product

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ProductHandler struct {
	productService IProductService
}

func NewProductHandler(productService IProductService) *ProductHandler {
	return &ProductHandler{productService}
}

// ShowProducts show
func (h *ProductHandler) ShowProducts(c *gin.Context) {
	products, err := h.productService.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, products)
}

// ShowProduct show
func (h *ProductHandler) ShowProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	product, err := h.productService.FindByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, product)
}

// CreateProduct create
func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var product Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	createdProduct, err := h.productService.Create(product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
		return
	}

	c.JSON(http.StatusCreated, createdProduct)
}

// UpdateProduct handles the request to update a product.
func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	// Extract product ID from the request params
	productID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	// Bind JSON request body to Product struct
	var updatedProduct Product
	if err := c.BindJSON(&updatedProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Call ProductService's Update method
	updatedProduct, err = h.productService.Update(productID, updatedProduct)
	if err != nil {
		// Handle the error, perhaps return an appropriate response to the client
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product"})
		return
	}

	// Handle the updated product, perhaps return it as a JSON response to the client
	c.JSON(http.StatusOK, updatedProduct)
}

// DestroyProduct remove a data record
func (h *ProductHandler) DestroyProduct(c *gin.Context) {
	productID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	// Call ProductService's Destroy method
	deletedProduct, err := h.productService.Destroy(productID)
	if err != nil {
		// Handle the error, perhaps return an appropriate response to the client
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	// Handle the deleted product, perhaps return it as a JSON response to the client
	c.JSON(http.StatusOK, deletedProduct)
}
