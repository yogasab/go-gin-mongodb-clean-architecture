package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-gin-mongodb-clean-architecture/app/dto"
	"go-gin-mongodb-clean-architecture/app/entities"
	campaign_image "go-gin-mongodb-clean-architecture/app/services/campaign-image"
	"go-gin-mongodb-clean-architecture/helpers"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

type campaignImageHandler struct {
	campaignImageService campaign_image.Service
}

func NewCampaignImageHandler(campaignImageService campaign_image.Service) *campaignImageHandler {
	return &campaignImageHandler{campaignImageService: campaignImageService}
}

func (h *campaignImageHandler) CreateCampaignImage(ctx *gin.Context) {
	var input dto.CreateCampaignImageInput
	err := ctx.ShouldBind(&input)
	if err != nil {
		response := helpers.APIResponse(http.StatusUnprocessableEntity, "failed", "Failed to process request", gin.H{"errors": err.Error()})
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	file, err := ctx.FormFile("campaign_image")
	if err != nil {
		response := helpers.APIResponse(http.StatusBadRequest, "failed", "Failed to upload image", gin.H{"errors": err.Error()})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	campaignImageID := primitive.NewObjectID()
	input.ID = campaignImageID

	user := ctx.MustGet("user").(entities.User)
	input.User = user

	fileName := fmt.Sprintf("assets/images/campaign/%s-%s", input.ID.Hex(), file.Filename)
	newCampaignImage, err := h.campaignImageService.CreateCampaignImage(input, fileName)
	if err != nil {
		response := helpers.APIResponse(http.StatusBadRequest, "failed", "Failed to create campaign campaign image", gin.H{"errors": err.Error()})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	err = ctx.SaveUploadedFile(file, fileName)
	if err != nil {
		response := helpers.APIResponse(http.StatusBadRequest, "failed", "Failed to upload image", gin.H{"errors": err.Error()})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := helpers.APIResponse(http.StatusCreated, "success", "Campaign image uploaded successfully", newCampaignImage)
	ctx.JSON(http.StatusCreated, response)
}
