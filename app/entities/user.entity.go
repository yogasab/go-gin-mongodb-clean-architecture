package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID             primitive.ObjectID `bson:"_id"`
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

func (u User) HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}
