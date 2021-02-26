package prices

import "github.com/samuelbeyeler/goarch-example/domain/product"

// ProductService Note here we are creating an interface that we need fulfilled, rather than pulling in the product service directly
type ProductService interface {
	GetAllByCompanyID(id string) []product.Product
}

// Service ...
type Service struct {
	repo     Repository
	products ProductService
}

// NewService accepts the interfaces as parameters.
func NewService(repo Repository, productService ProductService) Service {
	return Service{
		repo:     repo,
		products: productService,
	}
}

// GetPricesForAll ...
func (s Service) GetPricesForAll() []Price {
	var prices []Price
	products := s.products.GetAllByCompanyID("companyID")
	for _, prod := range products {
		prices = append(prices, s.repo.FindPriceByName(prod.Name))
	}
	return prices
}
