package main

import (
	"flag"
	"github.com/jinzhu/gorm"
	"solid-library-kata/internal/cli"
	"solid-library-kata/internal/repository"
)

func main() {
	f := cli.CommandFlags{}
	f.Action = flag.String("action", "", "action to do <list-books|add-book|add-user|borrow-book|get-borrowed-books|init")
	f.Book = flag.String("book", "", "book to add or borrow")
	f.Author = flag.String("author", "", "author of the book to add")
	f.User = flag.String("user", "", "username of the user")
	f.Role = flag.String("role", "", "role of the user to add")
	flag.Parse()

	db, err := gorm.Open("sqlite3", "/tmp/test.db")
	db.LogMode(true)
	if err != nil {
		panic("failed to connect database")
	}
	repo := repository.CreateRepository(db)

	cli.DispatchCommand(repo, f)

}
