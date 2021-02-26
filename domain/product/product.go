package product

// Repository ...
type Repository interface {
	FindByOwnerID(id string) []Product
}

// Product ...
type Product struct {
	Name string
}
