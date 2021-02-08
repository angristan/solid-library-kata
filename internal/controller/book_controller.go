package controller

import (
	"solid-library-kata/internal/model"
	"solid-library-kata/internal/output"
	"solid-library-kata/internal/repository"

	"github.com/olekukonko/tablewriter"
)

func AddBookToLibrary(title string, author string, userID string) *tablewriter.Table {
	user := repository.GetUser(userID)
	if !user.IsAdmin() {
		panic("user is not admin, not allowed")
	}

	if repository.BookExists(title) {
		panic("Book already exists")
	}

	repository.AddBook(model.Book{Title: title, AuthorName: author})
	book := repository.GetBook(title)

	return output.DisplayBook(book)
}

func ListBooks() *tablewriter.Table {
	// Guests can see the content of the library
	// So no user role check

	books := repository.GetAllBooks()

	return output.DisplayBooks(books)
}
