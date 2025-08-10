package models

import (
	"time"
)

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Username  string `gorm:"size:100;unique;not null"`
	Email     string `gorm:"size:255;unique;not null"`
	Password  string `gorm:"size:255;not null"`
	Role      string `gorm:"size:50;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
