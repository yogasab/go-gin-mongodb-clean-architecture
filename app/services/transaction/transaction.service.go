package transaction

import (
	"fmt"
	"go-gin-mongodb-clean-architecture/app/dto"
	"go-gin-mongodb-clean-architecture/app/entities"
	"go-gin-mongodb-clean-architecture/app/repositories/transaction"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Service interface {
	CreateTransaction(input dto.CreateTransactionInput) (string, error)
}

type service struct {
	transactionRepository transaction.Repository
}

func NewService(transactionRepository transaction.Repository) *service {
	return &service{transactionRepository}
}

func (s *service) CreateTransaction(input dto.CreateTransactionInput) (string, error) {
	transaction := entities.Transaction{}

	ID := primitive.NewObjectID()
	transaction.ID = ID
	transaction.User = input.User.ID
	campaignIDObj, _ := primitive.ObjectIDFromHex(input.Campaign)
	transaction.Campaign = campaignIDObj
	transaction.Code = fmt.Sprintf("TRX-%s", input.Campaign)
	transaction.Amount = input.Amount
	transaction.PaymentURL = ""
	transaction.Status = "PENDING"
	transaction.CreatedAt = time.Now()
	transaction.UpdatedAt = time.Now()

	newTransaction, err := s.transactionRepository.Create(transaction)
	if err != nil {
		return newTransaction, err
	}

	return newTransaction, nil
}
