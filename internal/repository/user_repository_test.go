package repository

import (
	"database/sql"
	"github.com/go-test/deep"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"regexp"
	"solid-library-kata/internal/model"
	"testing"
)

type Suite struct {
	suite.Suite
	DB         *gorm.DB
	mock       sqlmock.Sqlmock
	repository Repository
}

func (s *Suite) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)

	//s.DB, err = gorm.Open("postgres", db)
	s.DB, err = gorm.Open("sqlite3", db)
	require.NoError(s.T(), err)

	s.DB.LogMode(false)

	s.repository = CreateRepository(s.DB)
}

func (s *Suite) AfterTest(_, _ string) {
	require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func TestInit(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (s *Suite) TestGetUser() {
	var (
		ID       uint = 0
		username      = "testusername"
		role          = model.Admin
	)

	s.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "users" WHERE "users"."deleted_at" IS NULL AND ((username = ?)) ORDER BY "users"."id" ASC LIMIT 1`)).
		WithArgs(username).
		WillReturnRows(sqlmock.NewRows([]string{"id", "username", "role"}).
			AddRow(ID, username, role))

	res := s.repository.GetUser(username)

	//require.NoError(s.T(), err)
	require.Nil(s.T(), deep.Equal(model.User{ID: ID, Username: username, Role: role}, res))
}
