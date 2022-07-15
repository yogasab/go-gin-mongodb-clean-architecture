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

func (h *transactionHandler) GetTransactions(ctx *gin.Context) {
	var input dto.GetTransactionsInput

	user := ctx.MustGet("user").(entities.User)
	input.User = user

	transactions, err := h.transactionService.GetTransactions(input)
	if err != nil {
		response := helpers.APIResponse(http.StatusBadRequest, "failed", "Failed to get transactions", gin.H{"errors": err.Error()})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := helpers.APIResponse(http.StatusOK, "success", "Transactions fetched successfully", gin.H{"transactions": transactions})
	ctx.JSON(http.StatusOK, response)
}

func (h *transactionHandler) GetTransaction(ctx *gin.Context) {
	var input dto.GetTransactionInput

	user := ctx.MustGet("user").(entities.User)

	ID := ctx.Param("id")
	input.ID = ID
	input.User = user

	transaction, err := h.transactionService.GetTransaction(input)
	if err != nil {
		response := helpers.APIResponse(http.StatusBadRequest, "failed", "Failed to get transaction", gin.H{"errors": err.Error()})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := helpers.APIResponse(http.StatusOK, "success", "Transaction fetched successfully", gin.H{"transaction": transaction})
	ctx.JSON(http.StatusOK, response)
}
