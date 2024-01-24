package product

// ProductService defines the interface for product-related business logic.
type ProductService interface {
	FindAll() ([]Product, error)
	FindByID(ID int) (Product, error)
	Create(product Product) (Product, error)
}

// productService is the implementation of ProductService.
type productService struct {
	productRepository ProductRepository
}

// NewProductService creates a new instance of ProductService.
func NewProductService(repo ProductRepository) *productService {
	return &productService{repo}
}

// FindAll retrieves all products.
func (s *productService) FindAll() ([]Product, error) {
	return s.productRepository.FindAll()
}

// FindByID retrieves a product by its ID.
func (s *productService) FindByID(ID int) (Product, error) {
	return s.productRepository.FindByID(ID)
}

// Create adds a new product.
func (s *productService) Create(product Product) (Product, error) {
	return s.productRepository.Create(product)
}
