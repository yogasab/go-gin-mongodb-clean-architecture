package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID             primitive.ObjectID `bson:"id"`
	Name           string             `bson:"name"`
	Email          string             `bson:"email"`
	Password       string             `bson:"password"`
	AvatarFileName string             `bson:"avatar"`
	Role           string             `bson:"role"`
	Location       string             `bson:"location"`
	Occupation     string             `bson:"occupation"`
	CreatedAt      time.Time          `bson:"created_at"`
	UpdatedAt      time.Time          `bson:"updated_at"`
}
