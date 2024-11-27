package voucher

import (
	"otto/vouchers-project/models"

	"github.com/google/uuid"
)

type Service interface {
	CreateVoucher(name string, costInPoint int, brandID uuid.UUID) (*models.Voucher, error)
	GetVoucherByID(id string) (*models.Voucher, error)
	GetAllVoucherByBrand(id string) ([]models.Voucher, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) CreateVoucher(name string, costInPoint int, brandID uuid.UUID) (*models.Voucher, error) {
	voucher := &models.Voucher{
		ID:		uuid.New(),
		Name: name,
		CostInPoint: costInPoint,
		BrandID: brandID,
	}

	err := s.repo.CreateVoucher(voucher)
	return voucher, err
}

func (s *service) GetVoucherByID(id string) (*models.Voucher, error) {
	return s.repo.GetVoucherByID(id)
}

func (s *service) GetAllVoucherByBrand(brandID string) ([]models.Voucher, error) {
	return s.repo.GetAllVoucherByBrand(brandID)
}