package prices

import "github.com/samuelbeyeler/goarch-example/domain/product"

// Repository ...
type Repository interface {
	FindPriceByName(name string) Price
}

// Price For this example product is a direct dependency on the price struct
type Price struct {
	Price   float64
	Product product.Product
}
