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
	result := p.DB.Table("users").Where("username = ?", username)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		panic(result.Error)
	}

	return result.RowsAffected > 0
}

func (p *repo) GetUser(userID string) model.User {
	var user model.User

	result := p.DB.Where("username = ?", userID).First(&user)

	if result.Error != nil {
		panic(result.Error)
	}
	return user
}

func (p *repo) GetBorrowedBooks(user model.User) []model.Book {
	var books []model.Book
	result := p.DB.Table("books").Find(&books, "borrower_user_id = ?", user.ID)
	if result.Error != nil {
		panic(result.Error)
	}
	return books
}
