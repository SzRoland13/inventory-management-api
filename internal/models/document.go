package models

import (
	"time"
)

type Document struct {
	ID                uint      `gorm:"primaryKey"`
	Type              string    `gorm:"size:50;not null"`
	Number            string    `gorm:"size:100;unique;not null"`
	Date              time.Time `gorm:"type:date;not null"`
	PartnerID         uint      `gorm:"not null"`
	Partner           Partner   `gorm:"foreignKey:PartnerID"`
	CreatedByUserID   *uint
	CreatedByUser     *User `gorm:"foreignKey:CreatedByUserID"`
	WarehouseID       *uint
	Warehouse         *Warehouse `gorm:"foreignKey:WarehouseID"`
	TotalAmount       float64    `gorm:"type:numeric(15,2);default:0"`
	Status            string     `gorm:"size:50;not null"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
	CompletedAt       *time.Time
	RelatedDocumentID *uint
	RelatedDocument   *Document `gorm:"foreignKey:RelatedDocumentID"`
}
