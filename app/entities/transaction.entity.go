package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Transaction struct {
	ID         primitive.ObjectID `bson:"_id"`
	Campaign   primitive.ObjectID `bson:"campaign"`
	User       primitive.ObjectID `bson:"user"`
	Amount     int                `bson:"amount"`
	Status     string             `bson:"status"`
	Code       string             `bson:"code"`
	PaymentURL string             `bson:"payment_url"`
	Users      []User             `bson:"users"`
	Campaigns  []Campaign         `bson:"campaigns"`
	CreatedAt  time.Time          `bson:"created_at"`
	UpdatedAt  time.Time          `bson:"updated_at"`
}
