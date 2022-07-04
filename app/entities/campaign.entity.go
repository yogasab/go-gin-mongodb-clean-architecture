package entities

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Campaign struct {
	ID               primitive.ObjectID `bson:"_id"`
	User             primitive.ObjectID `bson:"user"`
	Title            string             `bson:"title"`
	ShortDescription string             `bson:"short_description"`
	Description      string             `bson:"description"`
	Perks            string             `bson:"perks"`
	BackerCount      int                `bson:"backer_count"`
	GoalAmount       int                `bson:"goal_amount"`
	CurrentAmount    int                `bson:"current_amount"`
	Slug             string             `bson:"slug"`
	CreatedAt        time.Time          `bson:"created_at"`
	UpdatedAt        time.Time          `bson:"updated_at"`
	Users            []User             `bson:"users"`
}
