package product

// IProductService defines the interface for product-related business logic.
type IProductService interface {
	FindAll() ([]Product, error)
	FindByID(ID int) (Product, error)
	Create(productRequest SProductRequest) (Product, error)
	Update(ID int, productRequest SProductRequest) (Product, error)
	Destroy(ID int) (Product, error)
}

// SProductService is the implementation of IProductService.
type SProductService struct {
	productRepository IProductRepository
}

// NewProductService creates a new instance of IProductService.
func NewProductService(productRepository IProductRepository) *SProductService {
	return &SProductService{productRepository}
}

// FindAll retrieves all products.
func (s *SProductService) FindAll() ([]Product, error) {
	products, err := s.productRepository.FindAll()
	return products, err
}

// FindByID retrieves a product by its ID.
func (s *SProductService) FindByID(ID int) (Product, error) {
	return s.productRepository.FindByID(ID)
}

// Create adds a new product.
func (s *SProductService) Create(productRequest SProductRequest) (Product, error) {
	qty, _ := productRequest.Qty.Float64()
	price, _ := productRequest.Price.Float64()

	product := Product{
		Name:        productRequest.Name,
		Description: productRequest.Description,
		Qty:         int(qty),
		Price:       int(price),
	}
	createdProduct, err := s.productRepository.Create(product)
	return createdProduct, err
}

// Update updates a product by ID.
func (s *SProductService) Update(ID int, productRequest SProductRequest) (Product, error) {
	// Check if the product with the given ID exists
	existingProduct, err := s.productRepository.FindByID(ID)
	if err != nil {
		return Product{}, err
	}

	// Perform any validation or business logic before updating (if needed)
	qty, _ := productRequest.Qty.Float64()
	price, _ := productRequest.Price.Float64()

	// Update the existing product with the new data
	existingProduct.Name = productRequest.Name
	existingProduct.Description = productRequest.Description
	existingProduct.Qty = int(qty)
	existingProduct.Price = int(price)
	// Update other fields as needed

	// Call the repository's Update method
	updatedProduct, err := s.productRepository.Update(existingProduct)
	return updatedProduct, err
}

// Destroy adds a new product.
func (s *SProductService) Destroy(ID int) (Product, error) {
	product, err := s.productRepository.FindByID(ID)
	deletedProduct, err := s.productRepository.Destroy(product)
	return deletedProduct, err
}
