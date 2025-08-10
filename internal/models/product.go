package models

import (
	"time"
)

type Product struct {
	ID          uint    `gorm:"primaryKey"`
	SKU         string  `gorm:"size:100;unique;not null"`
	Name        string  `gorm:"size:255;not null"`
	Description string  `gorm:"type:text"`
	Unit        string  `gorm:"size:50;not null"`
	NetPrice    float64 `gorm:"type:numeric(15,2);not null"`
	VATRate     float64 `gorm:"type:numeric(5,3);not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
