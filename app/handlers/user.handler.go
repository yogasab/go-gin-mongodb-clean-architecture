package handlers

import (
	"fmt"
	"go-gin-mongodb-clean-architecture/app/dto"
	"go-gin-mongodb-clean-architecture/app/entities"
	"go-gin-mongodb-clean-architecture/app/services/auth"
	"go-gin-mongodb-clean-architecture/app/services/user"
	"go-gin-mongodb-clean-architecture/helpers"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
	authService auth.Service
}

func NewUserHandler(userService user.Service, authService auth.Service) *userHandler {
	return &userHandler{userService: userService, authService: authService}
}

func (h *userHandler) GetAllUsers(ctx *gin.Context) {
	users, err := h.userService.GetAllUsers()
	if err != nil {
		response := helpers.APIResponse(http.StatusInternalServerError, "success", "Failed to process request", gin.H{"error": err.Error()})
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	usersFormatter := dto.FormatUsers(users)
	response := helpers.APIResponse(http.StatusOK, "success", "Users fetched successfully", usersFormatter)
	ctx.JSON(http.StatusOK, response)
}

func (h *userHandler) GetUserByID(ctx *gin.Context) {
	ID := ctx.Param("id")

	user, err := h.userService.GetUserByID(ID)
	if err != nil {
		response := helpers.APIResponse(http.StatusBadRequest, "failed", "Failed to process request", gin.H{"error": err.Error()})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	if user.Name == "" {
		response := helpers.APIResponse(http.StatusNotFound, "failed", "User not found", nil)
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	userFormatter := dto.FormatUser(user)
	response := helpers.APIResponse(http.StatusOK, "success", "User fetched successfully", userFormatter)
	ctx.JSON(http.StatusOK, response)
}

func (h *userHandler) CreateUser(ctx *gin.Context) {
	var input dto.CreateNewUserInput

	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		response := helpers.APIResponse(http.StatusUnprocessableEntity, "error", "Failed to process request", gin.H{"error": err.Error()})
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	var emailInput dto.CheckUserAvailabilityInput
	emailInput.Email = input.Email
	isNotRegistered, err := h.userService.CheckUserAvailability(emailInput)
	if !isNotRegistered {
		response := helpers.APIResponse(http.StatusBadRequest, "error", "Failed to create new user", "User with correspond email is already registered, please try another")
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	ID := primitive.NewObjectID()
	input.ID = ID
	jwtToken, err := h.authService.GenerateToken(ID.Hex())
	if err != nil {
		response := helpers.APIResponse(http.StatusBadRequest, "error", "Failed to authenticate user", gin.H{"error": err.Error()})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = h.userService.CreateUser(input, jwtToken)
	if err != nil {
		response := helpers.APIResponse(http.StatusBadRequest, "error", "Failed to create new user", gin.H{"error": err.Error()})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	response := helpers.APIResponse(http.StatusCreated, "success", "New user created successfully", jwtToken)
	ctx.JSON(http.StatusCreated, response)
}

func (h *userHandler) DeleteUserByID(ctx *gin.Context) {
	ID := ctx.Param("id")

	isDeleted, err := h.userService.DeleteUserByID(ID)
	if err != nil {
		response := helpers.APIResponse(http.StatusNotFound, "failed", err.Error(), gin.H{"error": err.Error()})
		ctx.JSON(http.StatusNotFound, response)
		return
	}
	response := helpers.APIResponse(http.StatusOK, "success", "User deleted successfully", isDeleted)
	ctx.JSON(http.StatusOK, response)
}

func (h *userHandler) UpdateUserByID(ctx *gin.Context) {
	var input dto.UpdateUserInput

	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		response := helpers.APIResponse(http.StatusUnprocessableEntity, "error", "Failed to process request", gin.H{"error": err.Error()})
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	ID := ctx.Param("id")
	input.ID = ID
	isUpdated, err := h.userService.UpdateUserByID(input)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response := helpers.APIResponse(http.StatusNotFound, "error", "Failed to update user", "User not found")
			ctx.JSON(http.StatusNotFound, response)
			return
		}
		response := helpers.APIResponse(http.StatusBadRequest, "error", "Failed to update user", gin.H{"error": err.Error()})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := helpers.APIResponse(http.StatusOK, "success", "User updated successfully", isUpdated)
	ctx.JSON(http.StatusOK, response)
}

func (h *userHandler) UploadUserAvatar(ctx *gin.Context) {
	var input dto.UpdloadUserAvatarInput

	err := ctx.ShouldBind(&input)
	if err != nil {
		response := helpers.APIResponse(http.StatusUnprocessableEntity, "failed", "Failed to process request", gin.H{"error": err.Error()})
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	user := ctx.MustGet("user").(entities.User)
	input.ID = user.ID.Hex()
	file, err := ctx.FormFile("avatar")
	if err != nil {
		response := helpers.APIResponse(http.StatusUnprocessableEntity, "failed", "Failed to upload image", gin.H{"error": err.Error()})
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	fileLocation := fmt.Sprintf("assets/images/user/%s-%s", input.ID, file.Filename)

	isUploaded, err := h.userService.UploadUserAvatar(input, fileLocation)
	if err != nil {
		response := helpers.APIResponse(http.StatusBadRequest, "failed", "Failed to upload image", gin.H{"error": err.Error()})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	if err = ctx.SaveUploadedFile(file, fileLocation); err != nil {
		response := helpers.APIResponse(http.StatusBadRequest, "failed", "Failed to upload image", gin.H{"error": err.Error()})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := helpers.APIResponse(http.StatusOK, "success", "Avatar uploaded successfully", isUploaded)
	ctx.JSON(http.StatusOK, response)
}

func (h *userHandler) LoginUser(ctx *gin.Context) {
	var input dto.LoginUserInput

	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		response := helpers.APIResponse(http.StatusUnprocessableEntity, "error", "Failed to process request", gin.H{"error": err.Error()})
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	loggedinUser, err := h.authService.LoginUser(input)
	if err != nil {
		response := helpers.APIResponse(http.StatusBadRequest, "error", "Failed to authenticate user", gin.H{"error": err.Error()})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	jwtToken, err := h.authService.GenerateToken(loggedinUser.ID.Hex())
	if err != nil {
		response := helpers.APIResponse(http.StatusUnauthorized, "failed", "Failed to authenticate user", gin.H{"error": err.Error()})
		ctx.JSON(http.StatusUnauthorized, response)
		return
	}

	userFormatter := dto.FormatUser(loggedinUser)
	response := helpers.APIResponse(http.StatusOK, "success", "User logged is successfully", gin.H{"user": userFormatter, "token": jwtToken})
	ctx.JSON(http.StatusOK, response)
}

func (h *userHandler) RegisterUser(ctx *gin.Context) {
	var input dto.CreateNewUserInput

	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		response := helpers.APIResponse(http.StatusUnprocessableEntity, "failed", "Failed to process request", gin.H{"error": err.Error()})
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	ID := primitive.NewObjectID()
	input.ID = ID
	jwtToken, err := h.authService.GenerateToken(ID.Hex())
	if err != nil {
		response := helpers.APIResponse(http.StatusBadRequest, "failed", "Failed to process request", gin.H{"error": err.Error()})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = h.authService.RegisterUser(input, jwtToken)
	if err != nil {
		response := helpers.APIResponse(http.StatusBadRequest, "failed", "Failed to authenticate user", gin.H{"error": err.Error()})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := helpers.APIResponse(http.StatusCreated, "success", "User registered successfully", jwtToken)
	ctx.JSON(http.StatusCreated, response)
}

func (h *userHandler) MyProfile(ctx *gin.Context) {
	user := ctx.MustGet("user").(entities.User)

	userFormatter := dto.FormatUser(user)
	response := helpers.APIResponse(http.StatusOK, "success", "Profile fetched successfully", userFormatter)
	ctx.JSON(http.StatusOK, response)
}
