package models

import (
	"time"
)

type DocumentLine struct {
	ID         uint     `gorm:"primaryKey"`
	DocumentID uint     `gorm:"not null;index"`
	Document   Document `gorm:"foreignKey:DocumentID"`
	ProductID  uint     `gorm:"not null"`
	Product    Product  `gorm:"foreignKey:ProductID"`
	Quantity   float64  `gorm:"type:numeric(15,2);not null"`
	UnitPrice  float64  `gorm:"type:numeric(15,2);not null"`
	TotalPrice float64  `gorm:"type:numeric(15,2);not null"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
