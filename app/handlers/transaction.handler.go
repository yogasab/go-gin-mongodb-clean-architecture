package handlers

import (
	"go-gin-mongodb-clean-architecture/app/dto"
	"go-gin-mongodb-clean-architecture/app/entities"
	"go-gin-mongodb-clean-architecture/app/services/transaction"
	"go-gin-mongodb-clean-architecture/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

type transactionHandler struct {
	transactionService transaction.Service
}

func NewTranscactionHandler(transactionService transaction.Service) *transactionHandler {
	return &transactionHandler{transactionService}
}

func (h *transactionHandler) CreateTransaction(ctx *gin.Context) {
	var input dto.CreateTransactionInput

	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		response := helpers.APIResponse(http.StatusUnprocessableEntity, "success", "Failed to process request", gin.H{"error": err.Error()})
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	user := ctx.MustGet("user").(entities.User)
	input.User = user

	newTransactionID, err := h.transactionService.CreateTransaction(input)
	if err != nil {
		response := helpers.APIResponse(http.StatusBadRequest, "error", "Failed to create new user", gin.H{"error": err.Error()})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := helpers.APIResponse(http.StatusCreated, "success", "New transaction created successfully", newTransactionID)
	ctx.JSON(http.StatusCreated, response)
}
