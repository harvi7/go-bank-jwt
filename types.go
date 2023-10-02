package main

import (
	"math/rand"
	"time"
)

type CreateAccountRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}
type Account struct {
	Id        int       `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Number    int64     `json:"number"`
	Balance   int64     `json:"balance"`
	CreatedAt time.Time `json:"createdAt"`
}

func NewAccount(firstname, lastName string) *Account {
	return &Account{
		FirstName: firstname,
		LastName:  lastName,
		Number:    int64(rand.Intn(100000)),
		Balance:   0,
		CreatedAt: time.Now().UTC(),
	}
}
