package model

import (
	"gorm.io/gorm"
	"time"
)

type Book struct {
	gorm.Model
	ID             uint `gorm:"primaryKey"`
	Title          string
	AuthorName     string
	BorrowerUserID uint `gorm:"DEFAULT:NULL"`
	BorrowedAt     time.Time
}
