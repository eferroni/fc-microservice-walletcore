package database

import (
	"database/sql"
	"testing"

	"github.com.br/eferroni/fc-ms-wallet/internal/entity"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
)

type ClientDbTestSuite struct {
	suite.Suite
	db *sql.DB
	clientDb *ClientDb
}

func (s *ClientDbTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Nil(err)
	s.db = db
	db.Exec("Create table clients (id varchar(255), name varchar(255), email varchar(255), created_at date)")
	s.clientDb = NewClientDB(db)
}

func (s *ClientDbTestSuite) TearDown() {
	defer s.db.Close()
	s.db.Exec("DROP TABLE clients")
}

func TestClientDbTestSuite(t *testing.T) {
	suite.Run(t, new(ClientDbTestSuite))
}

func (s *ClientDbTestSuite) TestSave() {
	client, _ := entity.NewClient("John Doe", "j@j.com")
	err := s.clientDb.Save(client)

	s.Nil(err)
}

func (s *ClientDbTestSuite) TestGet() {
	client, _ := entity.NewClient("John Doe", "j@j.com")
	s.clientDb.Save(client)

	clientDb, err := s.clientDb.Get(client.ID)
	s.Nil(err)
	s.Equal(client.ID, clientDb.ID)
	s.Equal(client.Name, clientDb.Name)
	s.Equal(client.Email, clientDb.Email)
}


