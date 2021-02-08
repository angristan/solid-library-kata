package cli

import (
	"solid-library-kata/internal/controller"
	"solid-library-kata/internal/output"
	"solid-library-kata/internal/util"
)

func DispatchCommand(f CommandFlags) {
	switch *f.Action {
	case "list-books":
		output.RenderOutput(controller.ListBooks())
	case "add-book":
		output.RenderOutput(controller.AddBookToLibrary(*f.Book, *f.Author, *f.User))
	case "add-user":
		output.RenderOutput(controller.AddUser(*f.User, *f.Role))
	case "init":
		util.InitDB()
	}
}
