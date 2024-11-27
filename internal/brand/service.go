package brand

import "otto/vouchers-project/models"

type Service interface {
	CreateBrand(brand *models.Brand) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) CreateBrand(brand *models.Brand) error {
	return s.repo.CreateBrand(brand)
}