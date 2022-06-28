package handlers

import (
	"go-gin-mongodb-clean-architecture/app/dto"
	"go-gin-mongodb-clean-architecture/app/services/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService: userService}
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

	newUser, err := h.userService.CreateUser(input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": "Failed to create new user", "status": "error", "errors": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"code": http.StatusCreated, "message": "New user created successfully", "status": "success", "data": newUser})
}
