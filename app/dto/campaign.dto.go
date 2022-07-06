package dto

import (
	"go-gin-mongodb-clean-architecture/app/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type CreateCampaignInput struct {
	User             primitive.ObjectID
	Title            string `json:"title" binding:"required"`
	ShortDescription string `json:"short_description" binding:"required"`
	Description      string `json:"description" binding:"required"`
	Perks            string `json:"perks" binding:"required"`
	GoalAmount       int    `json:"goal_amount" binding:"required"`
}

type UpdateCampaignInput struct {
	ID               string
	Title            string `json:"title" binding:"required"`
	ShortDescription string `json:"short_description" binding:"required"`
	Description      string `json:"description" binding:"required"`
	Perks            string `json:"perks" binding:"required"`
	GoalAmount       int    `json:"goal_amount" binding:"required"`
	User             entities.User
}

type UpdateCampaignBySlugInput struct {
	Slug             string
	Title            string `json:"title" binding:"required"`
	ShortDescription string `json:"short_description" binding:"required"`
	Description      string `json:"description" binding:"required"`
	Perks            string `json:"perks" binding:"required"`
	GoalAmount       int    `json:"goal_amount" binding:"required"`
	User             entities.User
}

type CampaignFormatter struct {
	ID               primitive.ObjectID `json:"id"`
	User             string             `json:"user"`
	Title            string             `json:"title"`
	ShortDescription string             `json:"short_description"`
	Description      string             `json:"description"`
	GoalAmount       int                `json:"goal_amount"`
	CurrentAmount    int                `json:"current_amount"`
	Slug             string             `json:"slug"`
	CreatedAt        time.Time          `json:"created_at"`
}

func FormatCampaign(campaign entities.Campaign) CampaignFormatter {
	campaignFormatter := CampaignFormatter{}
	campaignFormatter.ID = campaign.ID
	campaignFormatter.User = campaign.Users[0].Name
	campaignFormatter.Title = campaign.Title
	campaignFormatter.ShortDescription = campaign.ShortDescription
	campaignFormatter.Description = campaign.Description
	campaignFormatter.GoalAmount = campaign.GoalAmount
	campaignFormatter.Slug = campaign.Slug
	campaignFormatter.CreatedAt = campaign.CreatedAt

	return campaignFormatter
}

func FormatCampaigns(campaigns []entities.Campaign) []CampaignFormatter {
	var campaignsFormatter []CampaignFormatter

	for _, campaign := range campaigns {
		campaignFormatter := FormatCampaign(campaign)
		campaignsFormatter = append(campaignsFormatter, campaignFormatter)
	}

	return campaignsFormatter
}
