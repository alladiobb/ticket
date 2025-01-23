package database

import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"

	"github.com/alladiobb/ticket/internal/entity"
	"github.com/stretchr/testify/suite"
)

type AccountDBTestSuite struct {
	suite.Suite
	DB        *sql.DB
	accountDB *AccountDB
	client    *entity.Client
}

func (s *AccountDBTestSuite) SetupSuite() {
	// db, err := sql.Open("sqlite3", "file::memory:?cache=shared")
	db, err := sql.Open("sqlite3", ":memory:")
	s.Nil(err)
	s.DB = db

	db.Exec("CREATE TABLE clients (id varchar(255) PRIMARY KEY, name varchar(255), email varchar(255), created_at date)")
	db.Exec("CREATE TABLE accounts (id varchar(255) PRIMARY KEY, client_id varchar(255), balance decimal, created_at date)")
	s.accountDB = NewAccountDB(db)
	s.client, _ = entity.NewClient("John Doe", "j@j.com")
}

func (s *AccountDBTestSuite) TearDownSuite() {
	defer s.DB.Close()
	s.DB.Exec("DROP TABLE clients")
	s.DB.Exec("DROP TABLE accounts")
}

func TestAccountDBTestSuite(t *testing.T) {
	suite.Run(t, new(AccountDBTestSuite))
}

func (s *AccountDBTestSuite) TestSave() {
	account := entity.NewAccount(s.client)
	err := s.accountDB.Save(account)
	s.Nil(err)
}

func (s *AccountDBTestSuite) TestFindByID() {
	s.DB.Exec("INSERT INTO clients (id, name, email, created_at) VALUES (?, ?, ?, ?)",
		s.client.ID, s.client.Name, s.client.Email, s.client.CreatedAt,
	)

	account := entity.NewAccount(s.client)
	err := s.accountDB.Save(account)
	s.Nil(err)
	accountDB, err := s.accountDB.FindByID(account.ID)
	s.Nil(err)
	s.Equal(account.ID, accountDB.ID)
	s.Equal(account.Client.ID, accountDB.Client.ID)
	s.Equal(account.Balance, accountDB.Balance)
	s.Equal(account.Client.Name, accountDB.Client.Name)
	s.Equal(account.Client.Email, accountDB.Client.Email)
}
