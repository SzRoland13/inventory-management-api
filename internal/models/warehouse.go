package models

import (
	"time"
)

type Warehouse struct {
	ID              uint              `gorm:"primaryKey"`
	Name            string            `gorm:"size:150;unique;not null"`
	LocationID      uint              `gorm:"not null"`
	Location        WarehouseLocation `gorm:"foreignKey:LocationID"`
	CreatedByUserID *uint             // lehet null, pointerrel jelezve
	CreatedByUser   *User             `gorm:"foreignKey:CreatedByUserID"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
