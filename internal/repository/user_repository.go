package repository

import (
	"gorm.io/gorm"
	"solid-library-kata/internal/model"
)

func (p *repo) AddUser(user model.User) {
	result := p.DB.Create(&user)
	if result.Error != nil {
		panic(result.Error)
	}
}

func (p *repo) UserExists(username string) bool {
	var count int64

	result := p.DB.Table("users").Where("username = ?", username).Count(&count)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		panic(result.Error)
	}

	return count > 0
}

func (p *repo) GetUser(userID string) model.User {
	var user model.User

	result := p.DB.Where("username = ?", userID).First(&user)

	if result.Error != nil {
		panic(result.Error)
	}
	return user
}
