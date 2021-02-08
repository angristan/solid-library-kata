package main

import (
	"flag"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"solid-library-kata/internal/cli"
	"solid-library-kata/internal/repository"
)

func main() {
	f := cli.CommandFlags{}
	f.Action = flag.String("action", "", "action to do")
	f.Book = flag.String("book", "", "action to do")
	f.Author = flag.String("author", "", "action to do")
	f.User = flag.String("user", "", "action to do")
	f.Role = flag.String("role", "", "action to do")
	flag.Parse()

	db, err := gorm.Open(sqlite.Open("/tmp/test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	repo := repository.CreateRepository(db)

	cli.DispatchCommand(repo, f)

}
