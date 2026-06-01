package main

import (
	"math/rand"
	"time"
)

type CreateAccountRequest struct {
	FirstName string `json:"firstName"`
	LatsName string `json:"lastName"`
}

type Account struct{
	ID int `json:"id"`
	FirstName string `json:"firstName"`
	LatsName string `json:"lastName"`
	Number int64 `json:"number"`
	Balance int64 `json:"balance"`
	CreatedAt time.Time `json:"createdAt"`
}


func NewAccount(firstName,lastName string) *Account {
	return  &Account{
		FirstName: firstName,
		LatsName: lastName,
		Number: int64(rand.Intn(1000000000)),
		CreatedAt: time.Now().UTC(),
	}
}