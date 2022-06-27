package handlers

import (
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
