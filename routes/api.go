package routes

import (
	"go-gin-mongodb-clean-architecture/app/handlers"
	"go-gin-mongodb-clean-architecture/app/middlewares"
	campaignRepo "go-gin-mongodb-clean-architecture/app/repositories/campaign"
	campaignImageRepo "go-gin-mongodb-clean-architecture/app/repositories/campaign-image"
	transactionRepo "go-gin-mongodb-clean-architecture/app/repositories/transaction"
	userRepo "go-gin-mongodb-clean-architecture/app/repositories/user"
	"go-gin-mongodb-clean-architecture/app/services/auth"
	campaignServ "go-gin-mongodb-clean-architecture/app/services/campaign"
	campaignImageServ "go-gin-mongodb-clean-architecture/app/services/campaign-image"
	transactionServ "go-gin-mongodb-clean-architecture/app/services/transaction"
	userServ "go-gin-mongodb-clean-architecture/app/services/user"
	"go-gin-mongodb-clean-architecture/db"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine) {

	// User
	userCollection := db.GetCollection(db.DB, "users")
	userRepository := userRepo.NewUserRepository(userCollection)
	userService := userServ.NewService(userRepository)
	// Auth
	authService := auth.NewService(userService)
	// Campaign
	campaignCollection := db.GetCollection(db.DB, "campaigns")
	campaignRepository := campaignRepo.NewRepository(campaignCollection)
	campaignService := campaignServ.NewService(campaignRepository)
	// Campaign Image
	campaignImageCollection := db.GetCollection(db.DB, "campaign-images")
	campaignImageRepository := campaignImageRepo.NewRepository(campaignImageCollection)
	campaignImageService := campaignImageServ.NewService(campaignImageRepository, campaignRepository)
	// Transaction
	transactionCollection := db.GetCollection(db.DB, "transactions")
	transactionRepository := transactionRepo.NewRepository(transactionCollection)
	transactionService := transactionServ.NewService(transactionRepository)

	userAPIHandler := handlers.NewUserHandler(userService, authService)
	campaignHandler := handlers.NewCampaignHandler(campaignService)
	campaignImageHandler := handlers.NewCampaignImageHandler(campaignImageService)

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
		campaignAPIRouter.POST("/campaign-images", middlewares.AuthMiddleware(authService, userService), campaignImageHandler.CreateCampaignImage)
		campaignAPIRouter.PUT("/details/:slug", middlewares.AuthMiddleware(authService, userService), campaignHandler.UpdateCampaignBySlug)
		campaignAPIRouter.GET("/details/:slug", middlewares.AuthMiddleware(authService, userService), campaignHandler.GetCampaignBySlug)
	}
}
