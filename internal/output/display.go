package output

import (
	"github.com/olekukonko/tablewriter"
	"os"
	"solid-library-kata/internal/model"
)

func RenderOutput(table *tablewriter.Table) {
	table.Render()
}
func DisplayBooks(books []model.Book) *tablewriter.Table {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Title", "AuthorName"})

	for _, b := range books {
		table.Append([]string{b.Title, b.AuthorName})
	}
	return table
}

func DisplayBook(b model.Book) *tablewriter.Table {
	return DisplayBooks([]model.Book{
		b,
	})
}

func displayUsers(users []model.User) *tablewriter.Table {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Username", "Role"})

	for _, u := range users {
		table.Append([]string{u.Username, u.Role.GetString()})
	}
	return table
}

func DisplayUser(u model.User) *tablewriter.Table {
	return displayUsers([]model.User{
		u,
	})
}
