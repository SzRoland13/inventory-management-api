package models

import (
	"time"
)

type Partner struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"size:255;not null"`
	Type      string `gorm:"size:50;not null"`
	TaxNumber string `gorm:"size:50"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
