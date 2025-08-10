package models

import (
	"time"
)

type StockBalance struct {
	ID          uint      `gorm:"primaryKey"`
	WarehouseID uint      `gorm:"not null;index:idx_wh_product,unique"`
	Warehouse   Warehouse `gorm:"foreignKey:WarehouseID"`
	ProductID   uint      `gorm:"not null;index:idx_wh_product,unique"`
	Product     Product   `gorm:"foreignKey:ProductID"`
	Quantity    float64   `gorm:"type:numeric(15,2);default:0"`
	UpdatedAt   time.Time
}
