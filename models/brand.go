package models

import (
	"time"

	"github.com/google/uuid"
)


type Brand struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	Name      string    `json:"name"	gorm:"type:varchar(255);not null"`
	CreatedAt time.Time	`json:"-"`
	UpdatedAt time.Time `json:"-"`
}