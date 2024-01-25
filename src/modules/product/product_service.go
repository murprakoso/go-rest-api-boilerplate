package product

// IProductService defines the interface for product-related business logic.
type IProductService interface {
	FindAll() ([]Product, error)
	FindByID(ID int) (Product, error)
	Create(product Product) (Product, error)
	Update(ID int, product Product) (Product, error)
	Destroy(ID int) error
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

// Update adds a new product.
func (s *ProductService) Update(ID int, productRequest ProductRequest) (Product, error) {
	//product, _ := s.productRepository.FindByID(ID)
	return s.productRepository.Update(productRequest)
}
