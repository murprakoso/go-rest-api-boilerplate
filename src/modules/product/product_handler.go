package product

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type SProductHandler struct {
	productService IProductService
}

func NewProductHandler(productService IProductService) *SProductHandler {
	return &SProductHandler{productService}
}

// ShowProducts handles the HTTP GET request to retrieve all products.
func (h *SProductHandler) ShowProducts(c *gin.Context) {
	// Retrieve all products from the ProductService
	products, err := h.productService.FindAll()
	if err != nil {
		// If an error occurs, respond with an internal server error
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Map the product entities to a response format
	productListResponse := NewProductListResponseFromEntity(products)

	// Respond with the product list in the desired format
	c.JSON(http.StatusOK, gin.H{"data": productListResponse})
}

// ShowProduct handles the HTTP GET request to retrieve a single product by ID.
func (h *SProductHandler) ShowProduct(c *gin.Context) {
	// Extract product ID from the request parameters
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		// Respond with a Bad Request error if the product ID is invalid
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	// Retrieve the product details from the ProductService
	product, err := h.productService.FindByID(id)
	if err != nil {
		// Respond with a Not Found error if the product is not found
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	// Map the product entity to the response format
	productDetailResponse := NewProductDetailResponseFromEntity(product)

	// Respond with the product details in the desired format
	c.JSON(http.StatusOK, gin.H{"data": productDetailResponse})
}

// CreateProduct handles the HTTP POST request to create a new product.
func (h *SProductHandler) CreateProduct(c *gin.Context) {
	// Parse request body into SProductRequest struct
	var productRequest SProductRequest
	if err := c.ShouldBindJSON(&productRequest); err != nil {
		// If the request payload is invalid, respond with a Bad Request error
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Create the product using ProductService
	createdProduct, err := h.productService.Create(productRequest)
	if err != nil {
		// If there's an error while creating the product, respond with an Internal Server Error
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
		return
	}

	// Map the created product to a response format
	productResponse := NewProductDetailResponseFromEntity(createdProduct)

	// Respond with the created product details
	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    productResponse,
	})
}

// UpdateProduct handles the request to update a product.
func (h *SProductHandler) UpdateProduct(c *gin.Context) {
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

	// Call SProductService's Update method
	updatedProduct, err = h.productService.Update(productID, updatedProduct)
	if err != nil {
		// Handle the error, perhaps return an appropriate response to the client
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product"})
		return
	}

	// Map the created product to a response format
	productResponse := NewProductDetailResponseFromEntity(updatedProduct)

	// Handle the updated product, perhaps return it as a JSON response to the client
	// Respond with the updated product details
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    productResponse,
	})
}

// DestroyProduct remove a data record
func (h *SProductHandler) DestroyProduct(c *gin.Context) {
	productID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	// Call SProductService's Destroy method
	deletedProduct, err := h.productService.Destroy(productID)
	if err != nil {
		// Handle the error, perhaps return an appropriate response to the client
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	// Handle the deleted product, perhaps return it as a JSON response to the client
	c.JSON(http.StatusOK, deletedProduct)
}
