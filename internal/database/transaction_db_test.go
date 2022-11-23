package database

import (
	"database/sql"
	"testing"

	"github.com.br/eferroni/fc-ms-wallet/internal/entity"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
)

type TransactionDbTestSuite struct {
	suite.Suite
	db *sql.DB
	transactionDb *TransactionDb
	client1 *entity.Client
	client2 *entity.Client
	account1 *entity.Account
	account2 *entity.Account
}

func (s *TransactionDbTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Nil(err)
	s.db = db
	db.Exec("Create table clients (id varchar(255), name varchar(255), email varchar(255), created_at date)")
	db.Exec("Create table accounts (id varchar(255), client_id varchar(255), balance float, created_at date)")
	db.Exec("Create table transactions (id varchar(255), account_id_from varchar(255), account_id_to varchar(255), amount float, created_at date)")
	
	s.transactionDb = NewTransactionDB(db)
	s.client1, _ = entity.NewClient("John", "j@j.com")
	s.client2, _ = entity.NewClient("Mary", "m@j.com")
	s.account1 = entity.NewAccount(s.client1)
	s.account1.Balance = 1000
	s.account2 = entity.NewAccount(s.client2)
	s.account2.Balance = 1000
}

func (s *TransactionDbTestSuite) TearDown() {
	defer s.db.Close()
	s.db.Exec("DROP TABLE clients")
	s.db.Exec("DROP TABLE accounts")
	s.db.Exec("DROP TABLE transactions")
}

func TestTransactionDbTestSuite(t *testing.T) {
	suite.Run(t, new(TransactionDbTestSuite))
}

func (s *TransactionDbTestSuite) TestCreate() {
	transaction, err := entity.NewTransaction(s.account1, s.account2, 100)
	s.Nil(err)
	err = s.transactionDb.Create(transaction)
	s.Nil(err)
}



