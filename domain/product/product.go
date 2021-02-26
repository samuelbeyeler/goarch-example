// Package product is the most simple of packages, it does not depend on **any** external services.
package product

// Repository ...
type Repository interface {
	FindByOwnerID(id string) []Product
}

// Product ...
type Product struct {
	Name string
}
