package test

import (
	"go-gin-mongodb-clean-architecture/db"
	"testing"

	"go-gin-mongodb-clean-architecture/app/entities"
	transactionRepo "go-gin-mongodb-clean-architecture/app/repositories/transaction"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func TestFindByIDTransacton(t *testing.T) {
	objID, _ := primitive.ObjectIDFromHex("62ce383116ab6dcc787cc583")

	transaction, err := transactionRepository.FindByID(objID)
	if err != nil {
		t.Fatal(err.Error())
	}

	assert.NotNil(t, transaction)
	assert.NoError(t, err)
}

func TestCreateTransaction(t *testing.T) {
	objID := primitive.NewObjectID()
	objCampaignID, _ := primitive.ObjectIDFromHex("62c272d6c7cc6524da5a03e2")
	objUserID, _ := primitive.ObjectIDFromHex("62bd416dd08bdf54fe7ed518")

	newTransaction, err := transactionRepository.Create(entities.Transaction{ID: objID, Campaign: objCampaignID, User: objUserID, Amount: 120000, Status: "PENDING"})
	if err != nil {
		t.Fatalf(err.Error())
	}

	assert.NotNil(t, newTransaction)
	assert.NoError(t, err)
}
