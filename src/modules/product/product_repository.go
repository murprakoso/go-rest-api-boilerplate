package product

import (
	"gorm.io/gorm"
)

// IProductRepository defines the interface for interacting with the product data.
type IProductRepository interface {
	FindAll() ([]Product, error)
	FindByID(ID int) (Product, error)
	Create(product Product) (Product, error)
	Update(product Product) (Product, error)
	Destroy(product Product) (Product, error)
}

// SProductRepository is the implementation of IProductRepository.
type SProductRepository struct {
	db *gorm.DB
}

// NewProductRepository creates a new instance of IProductRepository.
func NewProductRepository(db *gorm.DB) *SProductRepository {
	return &SProductRepository{db}
}

// FindAll retrieves all products from the database.
func (r *SProductRepository) FindAll() ([]Product, error) {
	var products []Product
	err := r.db.Find(&products).Error
	return products, err
}

// FindByID retrieves a product by its ID from the database.
func (r *SProductRepository) FindByID(ID int) (Product, error) {
	var product Product
	err := r.db.First(&product, ID).Error
	return product, err
}

// Create adds a new product to the database.
func (r *SProductRepository) Create(product Product) (Product, error) {
	err := r.db.Create(&product).Error
	return product, err
}

// Update modifies an existing product in the database.
func (r *SProductRepository) Update(product Product) (Product, error) {
	err := r.db.Save(&product).Error
	return product, err
}

// Destroy removes a product from the database by its ID.
func (r *SProductRepository) Destroy(product Product) (Product, error) {
	err := r.db.Delete(&product).Error
	return product, err
}
