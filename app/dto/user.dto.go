package dto

import (
	"go-gin-mongodb-clean-architecture/app/entities"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CheckUserAvailabilityInput struct {
	Email string `json:"email" binding:"required,email"`
}

type CreateNewUserInput struct {
	ID         primitive.ObjectID
	Name       string `json:"name" binding:"required"`
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required"`
	Location   string `json:"location" binding:"required"`
	Role       string `json:"role" binding:"required"`
	Occupation string `json:"occupation" binding:"required"`
}

type UpdateUserInput struct {
	ID         string
	Name       string `json:"name,omitempty" binding:"required"`
	Email      string `json:"email,omitempty" binding:"required,email"`
	Password   string `json:"password,omitempty" binding:"required"`
	Location   string `json:"location,omitempty" binding:"required"`
	Occupation string `json:"occupation,omitempty" binding:"required"`
}

type UpdloadUserAvatarInput struct {
	ID string
}

type LoginUserInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UserFormatter struct {
	ID         primitive.ObjectID `json:"id"`
	Name       string             `json:"name"`
	Email      string             `json:"email"`
	Avatar     string             `json:"avatar"`
	Role       string             `json:"role"`
	Location   string             `json:"location"`
	Occupation string             `json:"occupation"`
	CreatedAt  time.Time          `json:"created_at"`
}

func FormatUser(user entities.User) UserFormatter {
	userFormatter := UserFormatter{}
	userFormatter.ID = user.ID
	userFormatter.Name = user.Name
	userFormatter.Email = user.Email
	userFormatter.Avatar = user.AvatarFileName
	userFormatter.Role = user.Role
	userFormatter.Location = user.Location
	userFormatter.Occupation = user.Occupation
	userFormatter.CreatedAt = user.CreatedAt

	return userFormatter
}

func FormatUsers(users []entities.User) []UserFormatter {
	var usersFormatter []UserFormatter

	for _, user := range users {
		formatUser := FormatUser(user)
		usersFormatter = append(usersFormatter, formatUser)
	}

	return usersFormatter
}
