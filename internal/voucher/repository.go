package voucher

import (
	"otto/vouchers-project/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	CreateVoucher(voucher *models.Voucher) error
	GetVoucherByID(id string) (*models.Voucher, error)
	GetAllByBrand(id string) ([]models.Voucher, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) CreateVoucher(voucher *models.Voucher) error {
	voucher.ID = uuid.New()
	return r.db.Create(voucher).Error
}

func (r *repository) GetVoucherByID(id string) (*models.Voucher, error) {
	var voucher models.Voucher
	if err := r.db.Where("id = ?", id).First(&voucher).Error; err != nil {
		return nil, err
	}

	return &voucher, nil
}

func (r *repository) GetAllByBrand(brandID string) ([]models.Voucher, error) {
	var vouchers []models.Voucher
	if err := r.db.Where("brand_id = ?", brandID).Find(&vouchers).Error; err != nil {
		return nil, err
	}

	return vouchers, nil
}