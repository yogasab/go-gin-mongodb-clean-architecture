package auth

import (
	"go-gin-mongodb-clean-architecture/app/dto"
	"go-gin-mongodb-clean-architecture/app/entities"
	"go-gin-mongodb-clean-architecture/app/services/user"

	"github.com/dgrijalva/jwt-go"
)

type Service interface {
	LoginUser(input dto.LoginUserInput) (entities.User, error)
	RegisterUser(input dto.CreateNewUserInput) (string, error)
}

type service struct {
	userService user.Service
}

var SECRET_KEY = []byte("s3cr3T_k3Y")

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

func (s *service) GenerateToken(UserID string) (string, error) {
	claims := jwt.MapClaims{}
	claims["user_id"] = UserID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	jwtToken, err := token.SignedString(SECRET_KEY)
	if err != nil {
		return jwtToken, err
	}

	return jwtToken, nil
}
