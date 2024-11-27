package brand

import (
	"otto/vouchers-project/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	CreateBrand(brand *models.Brand) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) CreateBrand(brand *models.Brand) error {
	brand.ID = uuid.New()
	return r.db.Create(brand).Error
}