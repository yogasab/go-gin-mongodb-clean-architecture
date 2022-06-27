package user

import (
	"go-gin-mongodb-clean-architecture/app/entities"
	"go-gin-mongodb-clean-architecture/app/repositories/user"
)

type Service interface {
	GetAllUsers() ([]entities.User, error)
}

type service struct {
	userRepository user.Reposisory
}

func NewService(userRepository user.Reposisory) *service {
	return &service{userRepository: userRepository}
}

func (s *service) GetAllUsers() ([]entities.User, error) {
	users, err := s.userRepository.FindAll()
	if err != nil {
		return users, err
	}

	return users, nil
}
