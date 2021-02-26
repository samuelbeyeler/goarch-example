// Package sql contains all of the repository implementations for the SQL data store.
// NOTE: if you wanted to move data stores, you would simply implement the interface in a new package (eg `mongo/product.go`)
package sql

import (
	"database/sql"
	"github.com/samuelbeyeler/goarch-example/domain/prices"
	"github.com/samuelbeyeler/goarch-example/domain/product"
)

// PriceRepo ...
type PriceRepo struct {
	client *sql.DB
}

// NewPriceRepo ...
func NewPriceRepo() PriceRepo {
	return PriceRepo{client: &sql.DB{}}
}

// FindPriceByName ...
func (r PriceRepo) FindPriceByName(name string) prices.Price {
	return prices.Price{
		Product: product.Product{},
	}
}
