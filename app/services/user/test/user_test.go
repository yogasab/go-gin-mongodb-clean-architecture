package test

import (
	"go-gin-mongodb-clean-architecture/app/dto"
	userRepo "go-gin-mongodb-clean-architecture/app/repositories/user"
	userServ "go-gin-mongodb-clean-architecture/app/services/user"
	"go-gin-mongodb-clean-architecture/db"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	userCollection = db.GetCollection(db.DB, "users")
	userRepository = userRepo.NewUserRepository(userCollection)
	userService    = userServ.NewService(userRepository)
)

func TestGetAllUsers(t *testing.T) {
	users, err := userService.GetAllUsers()
	totalUsers := len(users)
	if err != nil {
		t.Error(err)
	}

	assert.NotNil(t, users)
	assert.NoError(t, err)
	assert.Equal(t, len(users), totalUsers)
}

func TestGetUserByID(t *testing.T) {
	user, err := userService.GetUserByID("62bbc5f1a7dbcd9b551b7db5")
	if err != nil {
		t.Fatal(err)
	}

	assert.NotNil(t, user)
	assert.Equal(t, "62bbc5f1a7dbcd9b551b7db5", user.ID.Hex())
	assert.Equal(t, "Rodri", user.Name)
	assert.NoError(t, err)
}

func TestGetUserByEmail(t *testing.T) {
	user, err := userService.GetUserByEmail("rodrygo@gmail.com")
	if err != nil {
		t.Fatal(err)
	}

	assert.NotNil(t, user)
	assert.Equal(t, "rodrygo@gmail.com", user.Email)
	assert.Equal(t, "string", reflect.TypeOf(user.Email).String())
	assert.Equal(t, "Rodri", user.Name)
	assert.NoError(t, err)
}

func TestCheckUserAvailability(t *testing.T) {
	isAvailable, err := userService.CheckUserAvailability(dto.CheckUserAvailabilityInput{Email: "rodrygo@gmail.com"})
	if err != nil {
		t.Fatal(err)
	}

	assert.NotNil(t, isAvailable)
	assert.Equal(t, false, isAvailable)
	assert.Equal(t, "bool", reflect.Bool.String())
	assert.NoError(t, err)
}
