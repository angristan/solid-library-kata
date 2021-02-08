package repository

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"solid-library-kata/internal/model"
)

func AddBook(b model.Book) {
	db, err := gorm.Open(sqlite.Open("/tmp/test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	result := db.Create(&b)
	if result.Error != nil {
		panic(err)
	}
}

func GetBook(title string) model.Book {
	var book model.Book
	db, err := gorm.Open(sqlite.Open("/tmp/test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	result := db.Where("title = ?", title).First(&book)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		panic(result.Error)
	}

	return book
}

func BookExists(title string) bool {
	var count int64
	db, err := gorm.Open(sqlite.Open("/tmp/test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	result := db.Table("books").Where("title = ?", title).Count(&count)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		panic(result.Error)
	}

	return count > 0
}

func GetAllBooks() []model.Book {
	var books []model.Book
	db, err := gorm.Open(sqlite.Open("/tmp/test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	result := db.Find(&books)

	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		panic(result.Error)
	}
	return books
}

//func borrowBook(b model.Book) {
//
//}
