package user

import (
	"go-gin-mongodb-clean-architecture/app/entities"
	"go-gin-mongodb-clean-architecture/app/repositories/user"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Service interface {
	GetAllUsers() ([]entities.User, error)
	GetUserByID(ID string) (entities.User, error)
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

func (s *service) GetUserByID(ID string) (entities.User, error) {
	objID, _ := primitive.ObjectIDFromHex(ID)
	user, err := s.userRepository.FindByID(objID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			return user, nil
		}
		return user, err
	}

	return user, nil
}
