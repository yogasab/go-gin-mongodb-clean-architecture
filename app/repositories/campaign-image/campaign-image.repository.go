package campaign_image

import (
	"context"
	"fmt"
	"go-gin-mongodb-clean-architecture/app/entities"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type Repository interface {
	Create(campaignImage entities.CampaignImage) (string, error)
}

type repository struct {
	campaignImageCollection *mongo.Collection
}

func NewRepository(campaignImageCollection *mongo.Collection) *repository {
	return &repository{campaignImageCollection: campaignImageCollection}
}

func (r *repository) Create(campaignImage entities.CampaignImage) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := r.campaignImageCollection.InsertOne(ctx, campaignImage)
	if err != nil {
		return "", err
	}

	InsertedID := fmt.Sprintf("%s", result.InsertedID)
	return InsertedID, nil
}
