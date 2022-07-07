package entities

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type CampaignImage struct {
	ID        primitive.ObjectID `bson:"_id"`
	Campaign  primitive.ObjectID `bson:"campaign"`
	Filename  string             `bson:"filename"`
	IsPrimary bool               `bson:"is_primary"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}
