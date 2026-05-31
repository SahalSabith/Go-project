package main

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type Storage interface{
	CreateAccount(*Account) error
	DeleteAccount(int) error
	GetAccountbyid(int) (*Account, error)
	UpdateAccount(*Account) error
}


type PostgresStore struct{
	db *sql.DB
}

func NewPostgresStore()(*PostgresStore, error) {

}