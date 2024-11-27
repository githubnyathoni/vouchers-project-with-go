package models

import (
	"time"

	"github.com/google/uuid"
)

type Voucher struct {
	ID        	uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	BrandID   	uuid.UUID `json:"brand_id" gorm:"type:uuid;not null"`
	Name      	string    `json:"name"	gorm:"type:varchar(255);not null"`
	CostInPoint int     	`json:"cost_in_point"gorm:"type:int;not null"`
	CreatedAt 	time.Time	`json:"-"`
	UpdatedAt 	time.Time `json:"-"`
}