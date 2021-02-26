package product

// Service ...
type Service struct {
	repo Repository
}

// NewService ...
func NewService(repo Repository) Service {
	return Service{
		repo: repo,
	}
}

// GetAllByCompanyID ...
func (service Service) GetAllByCompanyID(id string) []Product {
	return service.repo.FindByOwnerID(id)
}
