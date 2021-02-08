package controller

import (
	"solid-library-kata/internal/model"
	"solid-library-kata/internal/output"
	"solid-library-kata/internal/repository"

	"github.com/olekukonko/tablewriter"
)

func AddBookToLibrary(repo repository.Repository, title string, author string, userID string) *tablewriter.Table {
	user := repo.GetUser(userID)
	if !user.IsAdmin() {
		panic("user is not admin, not allowed")
	}

	if repo.BookExists(title) {
		panic("Book already exists")
	}

	repo.AddBook(model.Book{Title: title, AuthorName: author})
	book := repo.GetBook(title)

	return output.DisplayBook(book)
}

func ListBooks(repo repository.Repository) *tablewriter.Table {
	// Guests can see the content of the library
	// So no user role check

	books := repo.GetAllBooks()

	return output.DisplayBooks(books)
}
