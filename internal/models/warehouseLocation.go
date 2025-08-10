package models

import (
	"time"
)

type WarehouseLocation struct {
	ID           uint   `gorm:"primaryKey"`
	AddressLine1 string `gorm:"size:255;not null"`
	AddressLine2 string `gorm:"size:255"`
	City         string `gorm:"size:100;not null"`
	Country      string `gorm:"size:2;not null"`
	Phone        string `gorm:"size:50"`
	Email        string `gorm:"size:255"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
