package routes

import (
	"github.com/gin-gonic/gin"
	"go-gin-mongodb-clean-architecture/app/dto"
	"go-gin-mongodb-clean-architecture/app/handlers"
	"go-gin-mongodb-clean-architecture/app/middlewares"
	campaignRepo "go-gin-mongodb-clean-architecture/app/repositories/campaign"
	userRepo "go-gin-mongodb-clean-architecture/app/repositories/user"
	"go-gin-mongodb-clean-architecture/app/services/auth"
	campaignServ "go-gin-mongodb-clean-architecture/app/services/campaign"
	userServ "go-gin-mongodb-clean-architecture/app/services/user"
	"go-gin-mongodb-clean-architecture/db"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

func InitializeRoutes(router *gin.Engine) {

	// user
	userCollection := db.GetCollection(db.DB, "users")
	userRepository := userRepo.NewUserRepository(userCollection)
	userService := userServ.NewService(userRepository)
	// Auth
	authService := auth.NewService(userService)
	// Campaign
	campaignCollection := db.GetCollection(db.DB, "campaigns")
	campaignRepository := campaignRepo.NewRepository(campaignCollection)
	campaignService := campaignServ.NewService(campaignRepository)

	userAPIHandler := handlers.NewUserHandler(userService, authService)

	input := dto.CreateCampaignInput{}
	objID, _ := primitive.ObjectIDFromHex("62bbc5f1a7dbcd9b551b7db5")
	input.User = objID
	input.Title = "Ayo Bantu Ukrain Lawan Russia"
	input.ShortDescription = "Ayo Bantu Ukrain Lawan Russia short short desc"
	input.Description = "Ayo Bantu Ukrain Lawan Russia short desc"
	input.Perks = "Menjadi dermawan, cerdas, dan masuk surga"
	input.GoalAmount = 120000000
	newCampaign, err := campaignService.Create(input)
	if err != nil {
		log.Fatalln(err.Error())
	}
	log.Fatalln(newCampaign)

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
