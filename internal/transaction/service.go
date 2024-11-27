package transaction

import (
	"errors"
	"otto/vouchers-project/models"

	"github.com/google/uuid"
)

type VoucherRepository interface {
	GetVoucherByID(id string) (*models.Voucher, error)
}

type Service interface {
	CreateTransaction(voucherID uuid.UUID, quantity int) (*models.Transaction, error)
	GetTransactionByID(id string) (*models.Transaction, error)
}

type service struct {
	repo 				Repository
	voucherRepo VoucherRepository
}

func NewService(repo Repository, voucherRepo VoucherRepository) Service {
	return &service{
		repo:        repo,
		voucherRepo: voucherRepo,
	}
}

func (s *service) CreateTransaction(voucherID uuid.UUID, quantity int) (*models.Transaction, error) {
	voucher, err := s.voucherRepo.GetVoucherByID(voucherID.String())
	if err != nil {
		return nil, errors.New("Voucher not found")
	}

	transaction := &models.Transaction{
		ID:	uuid.New(),
		VoucherID: voucherID,
		TotalPointsUsed: voucher.CostInPoint * quantity,
		Quantity: quantity,
		Status: "completed",
	}

	err = s.repo.CreateTransaction(transaction)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (s *service) GetTransactionByID(id string) (*models.Transaction, error) {
	return s.repo.GetTransactionByID(id)
}