package user

import (
	"context"
	"fmt"
	"go-gin-mongodb-clean-architecture/app/entities"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type CampaignRepository interface {
	Create(campaign entities.Campaign) (string, error)
}

type campaignRepository struct {
	campaignCollection *mongo.Collection
}

func NewRepository(campaignCollection *mongo.Collection) *campaignRepository {
	return &campaignRepository{campaignCollection: campaignCollection}
}

func (r *campaignRepository) Create(campaign entities.Campaign) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := r.campaignCollection.InsertOne(ctx, campaign)
	if err != nil {
		return "", err
	}

	insertedID := fmt.Sprintf("%s", result.InsertedID)
	return insertedID, nil
}
