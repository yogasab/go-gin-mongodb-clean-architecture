package transaction

import (
	"context"
	"fmt"
	"go-gin-mongodb-clean-architecture/app/entities"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	Create(transaction entities.Transaction) (string, error)
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
