package routes

import (
	"github.com/gin-gonic/gin"
	"go-gin-mongodb-clean-architecture/app/handlers"
	"go-gin-mongodb-clean-architecture/app/middlewares"
	userRepo "go-gin-mongodb-clean-architecture/app/repositories/user"
	"go-gin-mongodb-clean-architecture/app/services/auth"
	userServ "go-gin-mongodb-clean-architecture/app/services/user"
	"go-gin-mongodb-clean-architecture/db"
)

func InitializeRoutes(router *gin.Engine) {

	// user
	userCollection := db.GetCollection(db.DB, "users")
	userRepository := userRepo.NewUserRepository(userCollection)
	userService := userServ.NewService(userRepository)
	// auth
	authService := auth.NewService(userService)

	userAPIHandler := handlers.NewUserHandler(userService, authService)

	userAPIRouter := router.Group("/api/v1/users")
	{
		userAPIRouter.GET("/", middlewares.AuthMiddleware(authService, userService), userAPIHandler.GetAllUsers)
		userAPIRouter.POST("/", userAPIHandler.CreateUser)
		userAPIRouter.GET("/:id", userAPIHandler.GetUserByID)
		userAPIRouter.DELETE("/:id", userAPIHandler.DeleteUserByID)
		userAPIRouter.PUT("/:id", userAPIHandler.UpdateUserByID)
		userAPIRouter.POST("/avatars", userAPIHandler.UploadUserAvatar)
		userAPIRouter.GET("/profile", middlewares.AuthMiddleware(authService, userService), userAPIHandler.MyProfile)
	}

	authAPIRouter := router.Group("/api/v1/auth")
	{
		authAPIRouter.POST("/login", userAPIHandler.LoginUser)
		authAPIRouter.POST("/register", userAPIHandler.RegisterUser)
	}
}
