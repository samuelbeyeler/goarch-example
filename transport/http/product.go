package http

import (
	"github.com/samuelbeyeler/goarch-example/domain/product"
	"net/http"
)

// ProductService ...
type ProductService interface {
	GetAllByCompanyID(id string) []product.Product
}

// ProductHandler ...
type ProductHandler struct {
	products ProductService
}

// NewProductHandler ...
func NewProductHandler(products ProductService) ProductHandler {
	return ProductHandler{products: products}
}

// GetAllProducts ...
func (h ProductHandler) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	h.products.GetAllByCompanyID("companyID")
}
