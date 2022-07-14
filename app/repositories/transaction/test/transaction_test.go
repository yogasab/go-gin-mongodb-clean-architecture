package test

import (
	"go-gin-mongodb-clean-architecture/db"
	"testing"

	transactionRepo "go-gin-mongodb-clean-architecture/app/repositories/transaction"

	"github.com/stretchr/testify/assert"
)

var (
	transactionCollection = db.GetCollection(db.DB, "transactions")
	transactionRepository = transactionRepo.NewRepository(transactionCollection)
)

func TestFindAllTransaction(t *testing.T) {
	transactions, err := transactionRepository.FindAll()
	if err != nil {
		t.Fatal(err.Error())
	}

	assert.NotNil(t, transactions)
	assert.NoError(t, err)
}
