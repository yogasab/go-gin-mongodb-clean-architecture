package routes

import (
	"go-gin-mongodb-clean-architecture/app/handlers"
	userRepo "go-gin-mongodb-clean-architecture/app/repositories/user"
	userServ "go-gin-mongodb-clean-architecture/app/services/user"
	"go-gin-mongodb-clean-architecture/db"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine) {
	// user
	userCollection := db.GetCollection(db.DB, "users")
	userRepository := userRepo.NewUserRepository(userCollection)
	userService := userServ.NewService(userRepository)
	userAPIHandler := handlers.NewUserHandler(userService)

	userAPIRouter := router.Group("/api/v1/users")
	{
		userAPIRouter.GET("/", userAPIHandler.GetAllUsers)
		userAPIRouter.POST("/", userAPIHandler.CreateUser)
		userAPIRouter.GET("/:id", userAPIHandler.GetUserByID)
		userAPIRouter.DELETE("/:id", userAPIHandler.DeleteUserByID)
	}
}
