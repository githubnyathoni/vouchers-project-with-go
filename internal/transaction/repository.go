package transaction

import (
	"otto/vouchers-project/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	CreateTransaction(*models.Transaction) error
	GetTransactionByID(id string) (*models.Transaction, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) CreateTransaction(transaction *models.Transaction) error {
	transaction.ID = uuid.New()
	return r.db.Create(transaction).Error
}

func (r *repository) GetTransactionByID(id string) (*models.Transaction, error) {
	var transaction models.Transaction
	
	if err := r.db.Where("id = ?", id).First(&transaction).Error; err != nil {
		return nil, err
	}

	return &transaction, nil
}