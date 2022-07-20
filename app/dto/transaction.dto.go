package dto

import (
	"go-gin-mongodb-clean-architecture/app/entities"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateTransactionInput struct {
	Campaign string `json:"campaign"`
	Amount   int    `json:"amount"`
	User     entities.User
}

type GetTransactionsInput struct {
	User entities.User
}

type GetTransactionInput struct {
	ID   string
	User entities.User
}

type TransactionFormatter struct {
	ID         primitive.ObjectID `json:"_id"`
	Campaign   primitive.ObjectID `json:"campaign"`
	Amount     int                `json:"amount"`
	Status     string             `json:"status"`
	Code       string             `json:"code"`
	PaymentURL string             `json:"payment_url"`
	CreatedAt  time.Time          `json:"created_at"`
}

func FormatTransaction(transaction entities.Transaction) TransactionFormatter {
	transactionFormatter := TransactionFormatter{}
	transactionFormatter.ID = transaction.ID
	transactionFormatter.Campaign = transaction.Campaign
	transactionFormatter.Amount = transaction.Amount
	transactionFormatter.Status = transaction.Status
	transactionFormatter.Code = transaction.Code
	transactionFormatter.PaymentURL = transaction.PaymentURL
	transactionFormatter.CreatedAt = transaction.CreatedAt

	return transactionFormatter
}
