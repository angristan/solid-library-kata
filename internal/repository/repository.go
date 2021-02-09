package repository

import (
	"github.com/jinzhu/gorm"
	"solid-library-kata/internal/model"
)

type Repository interface {
	AddBook(b model.Book)
	GetBook(title string) model.Book
	BookExists(title string) bool
	GetAllBooks() []model.Book
	AddUser(user model.User)
	UserExists(username string) bool
	GetUser(userID string) model.User
}

type repo struct {
	DB *gorm.DB
}

func CreateRepository(db *gorm.DB) Repository {
	return &repo{
		DB: db,
	}
}
