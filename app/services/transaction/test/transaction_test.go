package test

import (
	"go-gin-mongodb-clean-architecture/app/dto"
	"go-gin-mongodb-clean-architecture/app/entities"
	transactionRepo "go-gin-mongodb-clean-architecture/app/repositories/transaction"
	transactionServ "go-gin-mongodb-clean-architecture/app/services/transaction"
	"go-gin-mongodb-clean-architecture/db"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	transactionCollection = db.GetCollection(db.DB, "transactions")
	transactionRepository = transactionRepo.NewRepository(transactionCollection)
	transactionService    = transactionServ.NewService(transactionRepository)
)

func TestTransactionService(t *testing.T) {
	objID, _ := primitive.ObjectIDFromHex("62ce383116ab6dcc787cc583")

	transactions, err := transactionService.GetTransactions(dto.GetTransactionsInput{User: entities.User{ID: objID, Role: "admin"}})

	if err != nil {
		t.Fatal(err)
	}

	assert.NotNil(t, transactions)
	assert.NoError(t, err)
}
