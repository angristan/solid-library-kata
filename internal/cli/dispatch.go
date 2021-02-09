package cli

import (
	"solid-library-kata/internal/controller"
	"solid-library-kata/internal/output"
	"solid-library-kata/internal/repository"
	"solid-library-kata/internal/util"
)

func DispatchCommand(repo repository.Repository, f CommandFlags) {
	switch *f.Action {
	case "list-books":
		output.RenderOutput(controller.ListBooks(repo))
	case "add-book":
		output.RenderOutput(controller.AddBookToLibrary(repo, *f.Book, *f.Author, *f.User))
	case "add-user":
		output.RenderOutput(controller.AddUser(repo, *f.User, *f.Role))
	case "borrow-book":
		output.RenderOutput(controller.BorrowBook(repo, *f.Book, *f.User))
	case "get-borrowed-books":
		output.RenderOutput(controller.GetBorrowedBooks(repo, *f.User))
	case "init":
		util.InitDB()
	}
}
