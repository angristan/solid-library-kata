package controller

import (
	"github.com/olekukonko/tablewriter"
	"solid-library-kata/internal/model"
	"solid-library-kata/internal/output"
	"solid-library-kata/internal/repository"
)

func AddUser(repo repository.Repository, userName string, roleName string) *tablewriter.Table {
	if repo.UserExists(userName) {
		panic("user already exists")
	}

	role := model.GetRole(roleName)
	user := model.User{Username: userName, Role: role}

	repo.AddUser(user)

	return output.DisplayUser(user)
}
