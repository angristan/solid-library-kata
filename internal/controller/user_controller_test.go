package controller

//
//import (
//	"database/sql"
//	"github.com/go-test/deep"
//	"github.com/jinzhu/gorm"
//	"github.com/stretchr/testify/require"
//	"github.com/stretchr/testify/suite"
//	"gopkg.in/DATA-DOG/go-sqlmock.v1"
//	"regexp"
//	"solid-library-kata/internal/model"
//	"solid-library-kata/internal/repository"
//	"testing"
//)
//
//type Suite struct {
//	suite.Suite
//	DB   *gorm.DB
//	mock sqlmock.Sqlmock
//	repository repository.Repository
//}
//
//func (s *Suite) SetupSuite() {
//	var (
//		db  *sql.DB
//		err error
//	)
//
//	db, s.mock, err = sqlmock.New()
//	require.NoError(s.T(), err)
//
//	//s.DB, err = gorm.Open("postgres", db)
//	s.DB, err = gorm.Open("sqlite3", db)
//	require.NoError(s.T(), err)
//	s.DB.LogMode(true)
//
//	s.repository = repository.CreateRepository(s.DB)
//}
//
//func (s *Suite) AfterTest(_, _ string) {
//	require.NoError(s.T(), s.mock.ExpectationsWereMet())
//}
//
//func TestInit(t *testing.T) {
//	suite.Run(t, new(Suite))
//}
//
//func (s *Suite) TestAddUser() {
//	var (
//		username      = "lol"
//		role          = "admin"
//		expectedRole = model.Admin
//	)
//	s.mock.ExpectBegin()

//	s.mock.ExpectQuery(regexp.QuoteMeta(
//		`SELECT count(*) FROM "users" WHERE (username = ?)`)).
//		WithArgs(username).
//		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(1))
//
//
//	user := AddUser(s.repository,username,role)
//
//	require.Nil(s.T(), deep.Equal(model.User{Username: username, Role: expectedRole}, user))
//}
