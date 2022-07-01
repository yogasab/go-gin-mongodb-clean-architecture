package handlers

import (
	"fmt"
	"go-gin-mongodb-clean-architecture/app/dto"
	"go-gin-mongodb-clean-architecture/app/entities"
	"go-gin-mongodb-clean-architecture/app/services/auth"
	"go-gin-mongodb-clean-architecture/app/services/user"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"

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
		ctx.JSON(http.StatusInternalServerError, gin.H{"errors": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "message": "Users fetched successfully", "status": "success", "users": users})
}

func (h *userHandler) GetUserByID(ctx *gin.Context) {
	ID := ctx.Param("id")

	user, err := h.userService.GetUserByID(ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"errors": err.Error()})
		return
	}
	if user.Name == "" {
		ctx.JSON(http.StatusNotFound, gin.H{"code": http.StatusNotFound, "message": "User not found", "status": "failed", "user": nil})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "message": "User fetched successfully", "status": "success", "user": user})
}

func (h *userHandler) CreateUser(ctx *gin.Context) {
	var input dto.CreateNewUserInput

	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": http.StatusUnprocessableEntity, "message": "Failed to process request", "status": "failed", "errors": err.Error()})
		return
	}

	var emailInput dto.CheckUserAvailabilityInput
	emailInput.Email = input.Email
	isNotRegistered, err := h.userService.CheckUserAvailability(emailInput)
	if !isNotRegistered {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": "Failed to create new user", "status": "error", "errors": "User with correspond email is already registered, please try another"})
		return
	}

	ID := primitive.NewObjectID()
	input.ID = ID
	jwtToken, err := h.authService.GenerateToken(ID.Hex())
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": "Failed to authenticate user", "status": "error", "errors": "Failed to authenticate user"})
		return
	}

	_, err = h.userService.CreateUser(input, jwtToken)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": "Failed to create new user", "status": "error", "errors": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"code": http.StatusCreated, "message": "New user created successfully", "status": "success", "data": jwtToken})
}

func (h *userHandler) DeleteUserByID(ctx *gin.Context) {
	ID := ctx.Param("id")

	isDeleted, err := h.userService.DeleteUserByID(ID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"code": http.StatusNotFound, "message": err.Error(), "status": "failed", "errors": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "message": "User deleted successfully", "status": "success", "is_deleted": isDeleted})
}

func (h *userHandler) UpdateUserByID(ctx *gin.Context) {
	var input dto.UpdateUserInput

	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": http.StatusUnprocessableEntity, "message": "Failed to process request", "status": "failed", "errors": err.Error()})
		return
	}

	ID := ctx.Param("id")
	input.ID = ID
	isUpdated, err := h.userService.UpdateUserByID(input)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			ctx.JSON(http.StatusNotFound, gin.H{"code": http.StatusNotFound, "message": "Failed to update user", "status": "error", "errors": "User not found"})
			return
		}
		ctx.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": "Failed to update user", "status": "error", "errors": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "message": "User updated successfully", "status": "success", "is_updated": isUpdated})
}

func (h *userHandler) UploadUserAvatar(ctx *gin.Context) {
	var input dto.UpdloadUserAvatarInput

	err := ctx.ShouldBind(&input)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": http.StatusUnprocessableEntity, "message": "Failed to process request", "status": "failed", "errors": err.Error()})
		return
	}

	user := ctx.MustGet("user").(entities.User)
	input.ID = user.ID.Hex()
	file, err := ctx.FormFile("avatar")
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": http.StatusUnprocessableEntity, "message": "Failed to upload image", "status": "failed", "errors": err.Error()})
		return
	}

	fileLocation := fmt.Sprintf("assets/images/%s-%s", input.ID, file.Filename)

	isUploaded, err := h.userService.UploadUserAvatar(input, fileLocation)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": "Failed to upload image", "status": "failed", "errors": err.Error()})
		return
	}

	if err = ctx.SaveUploadedFile(file, fileLocation); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": "Failed to upload image", "status": "failed", "errors": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "message": "Avatar uploaded successfully", "status": "success", "is_uploaded": isUploaded})
}

func (h *userHandler) LoginUser(ctx *gin.Context) {
	var input dto.LoginUserInput

	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": http.StatusUnprocessableEntity, "message": "Failed to process request", "status": "failed", "errors": err.Error()})
		return
	}

	loggedinUser, err := h.authService.LoginUser(input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": "Failed to authenticate user", "status": "error", "errors": err.Error()})
		return
	}

	jwtToken, err := h.authService.GenerateToken(loggedinUser.ID.Hex())
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"code": http.StatusOK, "message": "New user created successfully", "status": "success", "errors": err.Error()})
	}

	ctx.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "message": "New user created successfully", "status": "success", "user": loggedinUser, "token": jwtToken})
}

func (h *userHandler) RegisterUser(ctx *gin.Context) {
	var input dto.CreateNewUserInput

	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": http.StatusUnprocessableEntity, "message": "Failed to process request", "status": "failed", "errors": err.Error()})
		return
	}

	ID := primitive.NewObjectID()
	input.ID = ID
	jwtToken, err := h.authService.GenerateToken(ID.Hex())
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": "Failed to authenticate user", "status": "error", "errors": "Failed to authenticate user"})
		return
	}

	_, err = h.authService.RegisterUser(input, jwtToken)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": "Failed to authenticate user", "status": "error", "errors": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"code": http.StatusCreated, "message": "User registered successfully", "status": "success", "user": jwtToken})
}

func (h *userHandler) MyProfile(ctx *gin.Context) {
	user := ctx.MustGet("user").(entities.User)

	ctx.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "message": "Profile fetched successfully", "status": "success", "user": user})
}
