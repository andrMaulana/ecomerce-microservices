package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID          uint           `gorm:"primaryKey"`
	Name        string         `gorm:"not null"`
	Description string         `gorm:"type:text"`
	Price       float64        `gorm:"not null"`
	ImageURL    string         `gorm:"type:varchar(255)"`
	Inventory   int            `gorm:"not null"`
	CreatedAt   time.Time      `gorm:"not null"`
	UpdatedAt   time.Time      `gorm:"not null"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
