package database

import (
	"database/sql"
	"testing"

	"github.com.br/eferroni/fc-ms-wallet/internal/entity"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
)

type AccountDbTestSuite struct {
	suite.Suite
	db *sql.DB
	accountDb *AccountDb
	client *entity.Client
}

func (s *AccountDbTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Nil(err)
	s.db = db
	db.Exec("Create table clients (id varchar(255), name varchar(255), email varchar(255), created_at date)")
	db.Exec("Create table accounts (id varchar(255), client_id varchar(255), balance float, created_at date)")
	s.accountDb = NewAccountDB(db)
	s.client, _ = entity.NewClient("John", "j@j.com")
}

func (s *AccountDbTestSuite) TearDown() {
	defer s.db.Close()
	s.db.Exec("DROP TABLE clients")
	s.db.Exec("DROP TABLE accounts")
}

func TestAccountDbTestSuite(t *testing.T) {
	suite.Run(t, new(AccountDbTestSuite))
}

func (s *AccountDbTestSuite) TestSave() {
	account := entity.NewAccount(s.client)
	err := s.accountDb.Save(account)
	s.Nil(err)
}

func (s *AccountDbTestSuite) TestFindByID() {
	s.db.Exec("Insert into clients (id, name, email, created_at) values(?, ?, ?, ?)",
		s.client.ID, s.client.Name, s.client.Email, s.client.CreatedAt, 
	)
	account := entity.NewAccount(s.client)
	err := s.accountDb.Save(account)
	s.Nil(err)

	accountDb, err := s.accountDb.FindByID(account.ID)
	s.Nil(err)
	s.Equal(account.ID, accountDb.ID)
	s.Equal(account.Balance, accountDb.Balance)
	s.Equal(account.Client.ID, accountDb.Client.ID)
	s.Equal(account.Client.Name, accountDb.Client.Name)
	s.Equal(account.Client.Email, accountDb.Client.Email)
}


