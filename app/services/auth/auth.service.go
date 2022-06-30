package auth

import (
	"go-gin-mongodb-clean-architecture/app/dto"
	"go-gin-mongodb-clean-architecture/app/entities"
	"go-gin-mongodb-clean-architecture/app/services/user"
)

type Service interface {
	LoginUser(input dto.LoginUserInput) (entities.User, error)
	RegisterUser(input dto.CreateNewUserInput) (string, error)
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
	if err != nil {
		return user, err
	}

	if _, err = user.MatchedPassword(registeredUser.Password, password); err != nil {
		return user, err
	}

	return registeredUser, nil
}

func (s *service) RegisterUser(input dto.CreateNewUserInput) (string, error) {
	newUser, err := s.userService.CreateUser(input)
	if err != nil {
		return "", err
	}

	return newUser, nil
}
