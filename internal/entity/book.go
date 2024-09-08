package entity

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title       string `gorm:"not null"`
	Author      string `gorm:"not null"`
	IsAvailable bool   `gorm:"default:true"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
