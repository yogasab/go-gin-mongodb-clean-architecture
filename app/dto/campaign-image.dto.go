package dto

import "go.mongodb.org/mongo-driver/bson/primitive"

type CreateCampaignImageInput struct {
	ID        primitive.ObjectID
	Campaign  string `form:"campaign" binding:"required"`
	IsPrimary bool   `form:"is_primary"`
}
