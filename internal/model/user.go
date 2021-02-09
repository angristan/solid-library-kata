package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID            uint `gorm:"primaryKey"`
	Username      string
	Role          UserRole
	BorrowedBooks []Book `gorm:"foreignkey:BorrowerUserID;association_foreignkey:ID"`
}

type CreditCard struct {
	gorm.Model
	Number           string
	UserMemberNumber string
}

func (user *User) IsAdmin() bool {
	return user.Role == Admin
}
