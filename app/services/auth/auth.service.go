package auth

import (
	"errors"
	"go-gin-mongodb-clean-architecture/app/dto"
	"go-gin-mongodb-clean-architecture/app/entities"
	"go-gin-mongodb-clean-architecture/app/services/user"
)

type Service interface {
	LoginUser(input dto.LoginUserInput) (entities.User, error)
}

type service struct {
	userService user.Service
}

func NewService(userService user.Service) *service {
	return &service{userService: userService}
}

func (s *service) LoginUser(input dto.LoginUserInput) (entities.User, error) {
	var user entities.User
	email := input.Email
	password := input.Password

	registeredUser, err := s.userService.GetUserByEmail(email)
	if err == nil {
		return user, errors.New("User not registered, please register first")
	}

	isMatched, err := user.MatchedPassword(registeredUser.Password, password)
	if err != nil || !isMatched {
		return user, errors.New("Invalid credentials")
	}

	return registeredUser, nil
}
