package transaction

import (
	"context"
	"fmt"
	"go-gin-mongodb-clean-architecture/app/entities"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	Create(transaction entities.Transaction) (string, error)
	FindAll() ([]entities.Transaction, error)
}

type repository struct {
	transactionCollection *mongo.Collection
}

func NewRepository(transactionCollection *mongo.Collection) *repository {
	return &repository{transactionCollection: transactionCollection}
}

func (r *repository) Create(transaction entities.Transaction) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	newTransaction, err := r.transactionCollection.InsertOne(ctx, transaction)
	if err != nil {
		return "", err
	}

	InsertedID := fmt.Sprintf("%v", newTransaction.InsertedID)
	return InsertedID, nil
}

func (r *repository) FindAll() ([]entities.Transaction, error) {
	var transactions []entities.Transaction

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := r.transactionCollection.Find(ctx, bson.M{})
	if err != nil {
		return transactions, err
	}

	for cursor.Next(ctx) {
		var transaction entities.Transaction
		if err := cursor.Decode(&transaction); err != nil {
			return transactions, err
		}
		transactions = append(transactions, transaction)
	}

	return transactions, nil
}