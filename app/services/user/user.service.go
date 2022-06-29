package user

import (
	"errors"
	"go-gin-mongodb-clean-architecture/app/dto"
	"go-gin-mongodb-clean-architecture/app/entities"
	"go-gin-mongodb-clean-architecture/app/repositories/user"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Service interface {
	GetAllUsers() ([]entities.User, error)
	GetUserByID(ID string) (entities.User, error)
	GetUserByEmail(email string) (entities.User, error)
	CheckUserAvailability(input dto.CheckUserAvailabilityInput) (bool, error)
	CreateUser(input dto.CreateNewUserInput) (string, error)
	DeleteUserByID(ID string) (bool, error)
	UpdateUserByID(input dto.UpdateUserInput) (bool, error)
	UploadUserAvatar(input dto.UpdloadUserAvatarInput) (bool, error)
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

func (s *service) GetUserByEmail(email string) (entities.User, error) {
	var user entities.User
	user, err := s.userRepository.FindByEmail(email)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			return user, nil
		}
		return user, err
	}

	return user, nil
}

func (s *service) CheckUserAvailability(input dto.CheckUserAvailabilityInput) (bool, error) {
	isAvailable := false

	_, err := s.userRepository.FindByEmail(input.Email)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			isAvailable = true
			return isAvailable, nil
		}
		return isAvailable, err
	}

	return isAvailable, nil
}

func (s *service) CreateUser(input dto.CreateNewUserInput) (string, error) {
	user := entities.User{}
	user.ID = primitive.NewObjectID()
	user.Name = input.Name
	user.Email = input.Email
	user.Location = input.Location
	user.Occupation = input.Occupation
	user.Role = "user"
	user.AvatarFileName = ""
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	hashedPassword, _ := user.HashPassword(input.Password)
	user.Password = hashedPassword

	insertedID, err := s.userRepository.Create(user)
	if err != nil {
		return "", err
	}

	return insertedID, nil
}

func (s *service) DeleteUserByID(ID string) (bool, error) {
	objID, _ := primitive.ObjectIDFromHex(ID)
	result, err := s.userRepository.Delete(objID)
	if err != nil {
		return false, err
	}

	if result == 0 {
		return false, errors.New("User not found")
	}

	return true, nil
}

func (s *service) UpdateUserByID(input dto.UpdateUserInput) (bool, error) {
	isUpdated := true

	objID, _ := primitive.ObjectIDFromHex(input.ID)

	updatedUser, err := s.userRepository.FindByID(objID)
	if err != nil {
		isUpdated = false
		return isUpdated, err
	}
	updatedUser.Name = input.Name
	updatedUser.Email = input.Email
	updatedUser.Location = input.Location
	updatedUser.Occupation = input.Occupation

	hashedPassword, err := updatedUser.HashPassword(input.Password)
	if err != nil {
		isUpdated = false
		return isUpdated, err
	}
	updatedUser.Password = hashedPassword

	result, err := s.userRepository.UpdateByID(objID, updatedUser)
	if err != nil {
		isUpdated = false
		return isUpdated, err
	}

	if result == 0 {
		isUpdated = false
		return isUpdated, err
	}

	return isUpdated, nil
}

func (s *service) UploadUserAvatar(input dto.UpdloadUserAvatarInput) (bool, error) {
	isUploaded := true

	objID, _ := primitive.ObjectIDFromHex(input.ID)

	updatedUser, err := s.userRepository.FindByID(objID)
	if err != nil {
		isUploaded = false
		return isUploaded, err
	}

	updatedUser.AvatarFileName = input.Avatar
	result, err := s.userRepository.UpdateByID(objID, updatedUser)
	if err != nil {
		isUploaded = false
		return isUploaded, err
	}

	if result == 1 {
		isUploaded = true
	}
	return isUploaded, nil
}
