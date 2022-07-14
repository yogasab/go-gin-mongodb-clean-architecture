package dto

import "go-gin-mongodb-clean-architecture/app/entities"

type CreateTransactionInput struct {
	Campaign string `json:"campaign"`
	Amount   int    `json:"amount"`
	User     entities.User
}

type GetTransactionsInput struct {
	User entities.User
}
