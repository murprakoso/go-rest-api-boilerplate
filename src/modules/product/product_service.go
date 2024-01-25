package product

// IProductService defines the interface for product-related business logic.
type IProductService interface {
	FindAll() ([]Product, error)
	FindByID(ID int) (Product, error)
	Create(product Product) (Product, error)
	Update(ID int, product Product) (Product, error)
	Destroy(ID int) (Product, error)
}

// ProductService is the implementation of IProductService.
type ProductService struct {
	productRepository IProductRepository
}

// NewProductService creates a new instance of IProductService.
func NewProductService(productRepository IProductRepository) *ProductService {
	return &ProductService{productRepository}
}

// FindAll retrieves all products.
func (s *ProductService) FindAll() ([]Product, error) {
	return s.productRepository.FindAll()
}

// FindByID retrieves a product by its ID.
func (s *ProductService) FindByID(ID int) (Product, error) {
	return s.productRepository.FindByID(ID)
}

// Create adds a new product.
func (s *ProductService) Create(product Product) (Product, error) {
	return s.productRepository.Create(product)
}

// Update updates a product by ID.
func (s *ProductService) Update(ID int, product Product) (Product, error) {
	// Check if the product with the given ID exists
	existingProduct, err := s.productRepository.FindByID(ID)
	if err != nil {
		return Product{}, err
	}

	// Perform any validation or business logic before updating (if needed)

	// Update the existing product with the new data
	existingProduct.Name = product.Name
	existingProduct.Description = product.Description
	existingProduct.Qty = product.Qty
	existingProduct.Price = product.Price
	// Update other fields as needed

	// Call the repository's Update method
	updatedProduct, err := s.productRepository.Update(existingProduct)
	if err != nil {
		return Product{}, err
	}

	return updatedProduct, nil

	//product, _ := s.ProductRepository.FindByID(ID)
	//return s.productRepository.Update(product)
}

// Destroy adds a new product.
func (s *ProductService) Destroy(ID int) (Product, error) {
	product, err := s.productRepository.FindByID(ID)
	deletedProduct, err := s.productRepository.Destroy(product)
	return deletedProduct, err
}
