package database

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/wallison/fc-ms-wallet/internal/entity"
)

type TransactionDBTEstSuite struct {
	suite.Suite
	db            *sql.DB
	client        *entity.Client
	client2       *entity.Client
	accountFrom   *entity.Account
	accountTo     *entity.Account
	TransactionDB *TransactionDB
}

func (s *TransactionDBTEstSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Nil(err)
	s.db = db
	db.Exec("Create table clients (id varchar(255), name varchar(255), email varchar(255), created_at date)")
	db.Exec("Create table accounts (id varchar(255), client_id varchar(255), balance int, created_at date)")
	db.Exec("Create table transactions (id varchar(255), account_id_from varchar(255), account_id_to varchar(255), amount int, created_at date)")

	client, err := entity.NewClient("John", "j@j.com")
	s.Nil(err)
	s.client = client

	client2, err := entity.NewClient("John 2", "j@j2.com")
	s.Nil(err)
	s.client2 = client2

	accountFrom := entity.NewAccount(s.client)
	accountFrom.Balance = 1000
	s.accountFrom = accountFrom

	accountTo := entity.NewAccount(s.client2)
	accountTo.Balance = 1000
	s.accountTo = accountTo

	s.TransactionDB = NewTransactionDB(db)
}

func (s *TransactionDBTEstSuite) TearDownSuite() {
	defer s.db.Close()
	s.db.Exec("DROP TABLE clients")
	s.db.Exec("DROP TABLE accounts")
	s.db.Exec("DROP TABLE transactions")
}

func TestTransactionDBTestSuite(t *testing.T) {
	suite.Run(t, new(TransactionDBTEstSuite))
}

func (s *TransactionDBTEstSuite) TestCreate() {
	transaction, err := entity.NewTransaction(s.accountFrom, s.accountTo, 100)
	s.Nil(err)
	err = s.TransactionDB.Create(transaction)
	s.Nil(err)

}
