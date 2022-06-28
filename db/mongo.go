package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Client = ConnectDB()

func ConnectDB() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err.Error())
	}

	context, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err = client.Connect(context); err != nil {
		panic(err.Error())
	}
	if err = client.Ping(context, nil); err != nil {
		panic(err.Error())
	}

	log.Println("Database connected to MongoDB successfully")
	return client
}

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("go-gin-mongodb-clean-architecture").Collection(collectionName)

	return collection
}
