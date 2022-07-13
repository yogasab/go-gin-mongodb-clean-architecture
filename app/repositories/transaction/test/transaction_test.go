package test

import (
	"context"
	"go-gin-mongodb-clean-architecture/app/entities"
	"go-gin-mongodb-clean-architecture/db"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type repository struct {
	transactionCollection *mongo.Collection
}

func TestFindAllTransaction(t *testing.T) {
	transactionRepository := repository{
		transactionCollection: db.GetCollection(db.DB, "transactions"),
	}

	var transactions []entities.Transaction

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := transactionRepository.transactionCollection.Find(ctx, bson.M{})
	if err != nil {
		t.Fatal(err.Error())
	}

	for cursor.Next(ctx) {
		var transaction entities.Transaction
		if err := cursor.Decode(&transaction); err != nil {
			t.Fatal(err.Error())
		}
		transactions = append(transactions, transaction)
	}

	assert.NotNil(t, transactions)
	assert.NoError(t, err)
}
