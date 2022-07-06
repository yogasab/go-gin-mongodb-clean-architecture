package routes

import (
	"github.com/gin-gonic/gin"
	"go-gin-mongodb-clean-architecture/app/handlers"
	"go-gin-mongodb-clean-architecture/app/middlewares"
	campaignRepo "go-gin-mongodb-clean-architecture/app/repositories/campaign"
	userRepo "go-gin-mongodb-clean-architecture/app/repositories/user"
	"go-gin-mongodb-clean-architecture/app/services/auth"
	campaignServ "go-gin-mongodb-clean-architecture/app/services/campaign"
	userServ "go-gin-mongodb-clean-architecture/app/services/user"
	"go-gin-mongodb-clean-architecture/db"
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
	campaignHandler := handlers.NewCampaignHandler(campaignService)

	campaign, err := campaignService.GetCampaignBySlug("bantu-israel-62bd416dd08bdf54fe7ed511")
	if err != nil {
		log.Fatalln(err.Error())
	}
	log.Fatalln(campaign)

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

	campaignAPIRouter := router.Group("/api/v1/campaigns")
	{
		campaignAPIRouter.POST("/", middlewares.AuthMiddleware(authService, userService), campaignHandler.CreateCampaign)
		campaignAPIRouter.GET("/", middlewares.AuthMiddleware(authService, userService), campaignHandler.GetCampaigns)
		campaignAPIRouter.GET("/:id", middlewares.AuthMiddleware(authService, userService), campaignHandler.GetCampaign)
		campaignAPIRouter.PUT("/:id", middlewares.AuthMiddleware(authService, userService), campaignHandler.UpdateCampaignByID)
		campaignAPIRouter.PUT("/details/:slug", middlewares.AuthMiddleware(authService, userService), campaignHandler.UpdateCampaignBySlug)
	}
}
