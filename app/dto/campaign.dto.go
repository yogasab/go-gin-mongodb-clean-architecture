package dto

import "go.mongodb.org/mongo-driver/bson/primitive"

type CreateCampaignInput struct {
	User             primitive.ObjectID
	Title            string `json:"title" binding:"required"`
	ShortDescription string `json:"short_description" binding:"required"`
	Description      string `json:"description" binding:"required"`
	Perks            string `json:"perks" binding:"required"`
	GoalAmount       int    `json:"goal_amount" binding:"required"`
}
