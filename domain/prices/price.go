// Package prices is a 'complex' service, in that it depends on the `product` package in order to properly service price information.
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
