package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       uint `gorm:"primaryKey"`
	Username string
	Role     UserRole
}

func (user *User) IsAdmin() bool {
	return user.Role == Admin
}
