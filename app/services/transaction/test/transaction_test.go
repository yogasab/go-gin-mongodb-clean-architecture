package test

import (
	transactionRepo "go-gin-mongodb-clean-architecture/app/repositories/transaction"
	transactionServ "go-gin-mongodb-clean-architecture/app/services/transaction"
	"go-gin-mongodb-clean-architecture/db"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	transactionCollection = db.GetCollection(db.DB, "transactions")
	transactionRepository = transactionRepo.NewRepository(transactionCollection)
	transactionService    = transactionServ.NewService(transactionRepository)
)

func TestTransactionService(t *testing.T) {
	transactions, err := transactionService.GetTransactions()
	if err != nil {
		t.Fatal(err)
	}

	assert.NotNil(t, transactions)
	assert.NoError(t, err)
}
