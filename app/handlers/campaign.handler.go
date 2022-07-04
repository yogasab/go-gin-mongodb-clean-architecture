package handlers

import (
	"github.com/gin-gonic/gin"
	"go-gin-mongodb-clean-architecture/app/dto"
	"go-gin-mongodb-clean-architecture/app/entities"
	"go-gin-mongodb-clean-architecture/app/services/campaign"
	"go-gin-mongodb-clean-architecture/helpers"
	"net/http"
)

type campaignHandler struct {
	campaignService campaign.CampaignService
}

func NewCampaignHandler(campaignService campaign.CampaignService) *campaignHandler {
	return &campaignHandler{campaignService: campaignService}
}

func (h *campaignHandler) CreateCampaign(ctx *gin.Context) {
	var input dto.CreateCampaignInput
	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		response := helpers.APIResponse(http.StatusUnprocessableEntity, "failed", "Failed to process request", err.Error())
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	user := ctx.MustGet("user").(entities.User)
	input.User = user.ID

	newCampaign, err := h.campaignService.CreateCampaign(input)
	if err != nil {
		response := helpers.APIResponse(http.StatusBadRequest, "failed", "Failed to create campaign", err)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := helpers.APIResponse(http.StatusCreated, "success", "Campaign created successfully", newCampaign)
	ctx.JSON(http.StatusCreated, response)
}
