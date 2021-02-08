package repository

import (
	"gorm.io/gorm"
	"solid-library-kata/internal/model"
)

func (p *repo) AddBook(b model.Book) {
	result := p.DB.Create(&b)
	if result.Error != nil {
		panic(result.Error)
	}
}

func (p *repo) GetBook(title string) model.Book {
	var book model.Book

	result := p.DB.Where("title = ?", title).First(&book)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		panic(result.Error)
	}

	return book
}

func (p *repo) BookExists(title string) bool {
	var count int64

	result := p.DB.Table("books").Where("title = ?", title).Count(&count)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		panic(result.Error)
	}

	return count > 0
}

func (p *repo) GetAllBooks() []model.Book {
	var books []model.Book

	result := p.DB.Find(&books)

	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		panic(result.Error)
	}
	return books
}

//func (p *repo) borrowBook(b model.Book) {
//
//}
