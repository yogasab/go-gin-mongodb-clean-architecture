package test

import (
	userRepo "go-gin-mongodb-clean-architecture/app/repositories/user"
	userServ "go-gin-mongodb-clean-architecture/app/services/user"
	"go-gin-mongodb-clean-architecture/db"
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
	assert.Equal(t, "Rodri", user.Name)
	assert.NoError(t, err)
}
