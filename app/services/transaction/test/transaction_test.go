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

func TestGetTransactions(t *testing.T) {
	objID, _ := primitive.ObjectIDFromHex("62ce383116ab6dcc787cc583")

	transactions, err := transactionService.GetTransactions(dto.GetTransactionsInput{User: entities.User{ID: objID, Role: "admin"}})

	if err != nil {
		t.Fatal(err)
	}

	assert.NotNil(t, transactions)
	assert.NoError(t, err)
}

func TestGetTransaction(t *testing.T) {
	objID, _ := primitive.ObjectIDFromHex("62bd416dd08bdf54fe7ed518")

	transaction, err := transactionService.GetTransaction(dto.GetTransactionInput{ID: "62ce383116ab6dcc787cc583", User: entities.User{ID: objID}})

	if err != nil {
		t.Fatal(err)
	}

	assert.NotNil(t, transaction)
	assert.NoError(t, err)
}

func TestCreateTransaction(t *testing.T) {
	objID, _ := primitive.ObjectIDFromHex("62bd416dd08bdf54fe7ed518")

	newTransaction, err := transactionService.CreateTransaction(dto.CreateTransactionInput{Campaign: "62c272d6c7cc6524da5a03e2", User: entities.User{ID: objID}, Amount: 120000})
	if err != nil {
		t.Fatalf(err.Error())
	}

	assert.NotNil(t, newTransaction)
	assert.NoError(t, err)
}

func TestGetUserTransactions(t *testing.T) {
	transactions, err := transactionService.GetUserTransaction("62bbc5f1a7dbcd9b551b7db5")
	if err != nil {
		t.Fatal(err.Error())
	}

	assert.NotNil(t, transactions)
	assert.NoError(t, err)
}
