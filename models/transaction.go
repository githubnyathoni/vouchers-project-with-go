package models

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID        			uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	VoucherID   		uuid.UUID `json:"voucher_id" gorm:"type:uuid;not null"`
	TotalPointsUsed	int				`json:"total_points_used" gorm:"type:int;not null"`
	Quantity				int				`json:"quantity" gorm:"type:int;not null"`
	Status					string		`json:"status" gorm:"type:varchar(255);not null"`
	CreatedAt 			time.Time	`json:"-"`
	UpdatedAt 			time.Time `json:"-"`
}