package util

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"solid-library-kata/internal/model"
)

func InitDB() {
	db, err := gorm.Open(sqlite.Open("/tmp/test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// Migrate the schema
	db.AutoMigrate(&model.Book{})
	db.AutoMigrate(&model.User{})

	// Create
	db.Create(&model.Book{Title: "Software Engineering", AuthorName: "Rui"})
	db.Create(&model.User{Username: "stan", Role: model.Admin})
}
