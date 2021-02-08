package repository

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"solid-library-kata/internal/model"
)

func AddUser(user model.User) {
	db, err := gorm.Open(sqlite.Open("/tmp/test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	result := db.Create(&user)
	if result.Error != nil {
		panic(err)
	}
}

func UserExists(username string) bool {
	var count int64
	db, err := gorm.Open(sqlite.Open("/tmp/test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	result := db.Table("users").Where("username = ?", username).Count(&count)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		panic(result.Error)
	}

	return count > 0
}

func GetUser(userID string) model.User {
	var user model.User
	db, err := gorm.Open(sqlite.Open("/tmp/test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	result := db.Where("username = ?", userID).First(&user)

	if result.Error != nil {
		panic(err)
	}
	return user
}
