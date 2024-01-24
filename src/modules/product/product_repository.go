package product

import (
	"gorm.io/gorm"
)

// ProductRepository defines the interface for interacting with the product data.
type ProductRepository interface {
	FindAll() ([]Product, error)
	FindByID(ID int) (Product, error)
	Create(product Product) (Product, error)
}

// productRepository is the implementation of ProductRepository.
type productRepository struct {
	db *gorm.DB
}

// NewProductRepository creates a new instance of ProductRepository.
func NewProductRepository(db *gorm.DB) *productRepository {
	return &productRepository{db}
}

// FindAll retrieves all products from the database.
func (r *productRepository) FindAll() ([]Product, error) {
	var products []Product
	if err := r.db.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

// FindByID retrieves a product by its ID from the database.
func (r *productRepository) FindByID(ID int) (Product, error) {
	var product Product
	if err := r.db.First(&product, ID).Error; err != nil {
		return Product{}, err
	}
	return product, nil
}

// Create adds a new product to the database.
func (r *productRepository) Create(product Product) (Product, error) {
	if err := r.db.Create(&product).Error; err != nil {
		return Product{}, err
	}
	return product, nil
}
