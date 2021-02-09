package model

import (
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
)

type Suite struct {
	suite.Suite
}

func TestInit(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (s *Suite) TestGetUser() {
	admin := User{
		ID:            0,
		Username:      "stan",
		Role:          0,
		BorrowedBooks: nil,
	}
	member := User{
		ID:            0,
		Username:      "stan",
		Role:          1,
		BorrowedBooks: nil,
	}
	guest := User{
		ID:            0,
		Username:      "stan",
		Role:          2,
		BorrowedBooks: nil,
	}
	require.True(s.T(), admin.IsAdmin())
	require.False(s.T(), member.IsAdmin())
	require.False(s.T(), guest.IsAdmin())
}
