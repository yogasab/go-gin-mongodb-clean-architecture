package campaign

import (
	"context"
	"fmt"
	"go-gin-mongodb-clean-architecture/app/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type Repository interface {
	Create(campaign entities.Campaign) (string, error)
	FindBySlug(slug string) (entities.Campaign, error)
}

type repository struct {
	campaignCollection *mongo.Collection
}

func NewRepository(campaignCollection *mongo.Collection) *repository {
	return &repository{campaignCollection: campaignCollection}
}

func (r *repository) Create(campaign entities.Campaign) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := r.campaignCollection.InsertOne(ctx, campaign)
	if err != nil {
		return "", err
	}

	insertedID := fmt.Sprintf("%s", result.InsertedID)
	return insertedID, nil
}

func (r *repository) FindBySlug(slug string) (entities.Campaign, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var campaign entities.Campaign
	filter := bson.M{"slug": slug}
	err := r.campaignCollection.FindOne(ctx, filter).Decode(&campaign)
	if err != nil {
		return campaign, err
	}

	return campaign, nil
}
