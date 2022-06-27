package routes

import (
	userRepo "go-gin-mongodb-clean-architecture/app/repositories/user"
	userServ "go-gin-mongodb-clean-architecture/app/services/user"
	"go-gin-mongodb-clean-architecture/db"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine) {
	// user
	userCollection := db.GetCollection(db.DB, "users")
	userRepository := userRepo.NewUserRepository(userCollection)
	_ = userServ.NewService(userRepository)
}
