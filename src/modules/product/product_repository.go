package product

import (
	"gorm.io/gorm"
)

// IProductRepository defines the interface for interacting with the product data.
type IProductRepository interface {
	FindAll() ([]Product, error)
	FindByID(ID int) (Product, error)
	Create(productRequest ProductRequest) (Product, error)
	Update(productRequest ProductRequest) (Product, error)
	Destroy(ID int) error
}

// ProductRepository is the implementation of IProductRepository.
type ProductRepository struct {
	db *gorm.DB
}

// NewProductRepository creates a new instance of IProductRepository.
func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db}
}

// FindAll retrieves all products from the database.
func (r *ProductRepository) FindAll() ([]Product, error) {
	var products []Product
	if err := r.db.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

// FindByID retrieves a product by its ID from the database.
func (r *ProductRepository) FindByID(ID int) (Product, error) {
	var product Product
	if err := r.db.First(&product, ID).Error; err != nil {
		return Product{}, err
	}
	return product, nil
}

// Create adds a new product to the database.
func (r *ProductRepository) Create(productRequest ProductRequest) (Product, error) {
	product, err := r.db.Create(&productRequest).Error
	//if err := r.db.Create(&productRequest).Error; err != nil {
	//	return Product{}, err
	//}
	return product, nil
}

// Update modifies an existing product in the database.
func (r *ProductRepository) Update(product Product) (Product, error) {
	if err := r.db.Save(&product).Error; err != nil {
		return Product{}, err
	}
	return product, nil
}

// Destroy removes a product from the database by its ID.
func (r *ProductRepository) Destroy(ID int) error {
	if err := r.db.Delete(&Product{}, ID).Error; err != nil {
		return err
	}
	return nil
}
