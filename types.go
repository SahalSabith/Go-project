package main

import "math/rand"

type Account struct{
	ID int `json:id`
	FirstName string `json:firstName`
	LatsName string `json:lastName`
	Number int64 `json:number`
	Balance int64 `json:balance`
}


func NewAccount(firstName,lastName string) *Account {
	return  &Account{
		ID: rand.Intn(10000),
		FirstName: firstName,
		LatsName: lastName,
		Number: int64(rand.Intn(1000000000)),
	}
}