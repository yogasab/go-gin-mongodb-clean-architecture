package transaction

import (
	"context"
	"fmt"
	"go-gin-mongodb-clean-architecture/app/entities"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	Create(transaction entities.Transaction) (string, error)
	FindAll() ([]entities.Transaction, error)
	FindByID(ID primitive.ObjectID) (entities.Transaction, error)
	FindByUserID(UserID primitive.ObjectID) ([]entities.Transaction, error)
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

func (r *repository) FindByID(ID primitive.ObjectID) (entities.Transaction, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var transaction entities.Transaction
	filter := bson.M{"_id": ID}
	err := r.transactionCollection.FindOne(ctx, filter).Decode(&transaction)
	if err != nil {
		return transaction, err
	}

	return transaction, nil
}

func (r *repository) FindByUserID(UserID primitive.ObjectID) ([]entities.Transaction, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	searchFilter := bson.M{"user": UserID}
	aggSearch := bson.M{"$match": searchFilter}
	aggPopulate := bson.M{
		"$lookup": bson.M{
			"from":         "users",
			"localField":   "user",
			"foreignField": "_id",
			"as":           "users",
		},
	}

	var transactions []entities.Transaction
	cursor, err := r.transactionCollection.Aggregate(ctx, []bson.M{aggSearch, aggPopulate})
	if err != nil {
		return transactions, err
	}

	for cursor.Next(ctx) {
		var transaction entities.Transaction
		err := cursor.Decode(&transaction)
		if err != nil {
			return transactions, err
		}
		transactions = append(transactions, transaction)
	}

	return transactions, nil
}
