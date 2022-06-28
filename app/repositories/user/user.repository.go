package user

import (
	"context"
	"fmt"
	"go-gin-mongodb-clean-architecture/app/entities"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Reposisory interface {
	FindAll() ([]entities.User, error)
	FindByID(ID primitive.ObjectID) (entities.User, error)
	FindByEmail(email string) (entities.User, error)
	Create(user entities.User) (string, error)
}

type repository struct {
	userCollection *mongo.Collection
}

func NewUserRepository(userCollection *mongo.Collection) *repository {
	return &repository{userCollection: userCollection}
}

func (r *repository) FindAll() ([]entities.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var users []entities.User
	cursor, err := r.userCollection.Find(ctx, bson.M{})
	if err != nil {
		return users, err
	}

	for cursor.Next(ctx) {
		var user entities.User
		if err = cursor.Decode(&user); err != nil {
			return users, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (r *repository) FindByID(ID primitive.ObjectID) (entities.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var user entities.User
	filter := bson.M{"id": ID}
	if err := r.userCollection.FindOne(ctx, filter).Decode(&user); err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) FindByEmail(email string) (entities.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var user entities.User
	filter := bson.M{"email": email}
	if err := r.userCollection.FindOne(ctx, filter).Decode(&user); err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) Create(user entities.User) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := r.userCollection.InsertOne(ctx, user)
	if err != nil {
		return "", err
	}

	insertedID := fmt.Sprintf("%s", result.InsertedID)
	return insertedID, nil
}
