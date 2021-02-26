package sql

import (
	"database/sql"
	"github.com/samuelbeyeler/goarch-example/domain/product"
)

// ProductRepo ...
type ProductRepo struct {
	client *sql.DB
}

// NewProductRepo ...
func NewProductRepo() ProductRepo {
	return ProductRepo{client: &sql.DB{}}
}

// FindByOwnerID ...
func (repo ProductRepo) FindByOwnerID(id string) []product.Product {
	return []product.Product{}
}
