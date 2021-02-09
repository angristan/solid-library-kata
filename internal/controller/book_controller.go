package controller

import (
	"solid-library-kata/internal/model"
	"solid-library-kata/internal/output"
	"solid-library-kata/internal/repository"
	"time"

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

func BorrowBook(repo repository.Repository, bookTitle string, userName string) *tablewriter.Table {
	book := repo.GetBook(bookTitle)
	user := repo.GetUser(userName)
	borrowedBooks := repo.GetBorrowedBooks(user)

	if len(borrowedBooks) >= 3 {
		panic("You can't borrow more than 3 books")
	}

	for _, book := range borrowedBooks {
		if book.BorrowedAt.Before(time.Now().Add(4 * 7 * 24 * time.Hour)) {
			panic("You already have a book due: " + book.Title)
		}
	}

	book = repo.BorrowBook(book, user)

	return output.DisplayBook(book)
}

func GetBorrowedBooks(repo repository.Repository, userName string) *tablewriter.Table {
	user := repo.GetUser(userName)
	borrowedBooks := repo.GetBorrowedBooks(user)

	return output.DisplayBooks(borrowedBooks)
}
