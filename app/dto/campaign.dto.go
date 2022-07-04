package dto

import "go.mongodb.org/mongo-driver/bson/primitive"

type CreateCampaignInput struct {
	User             primitive.ObjectID
	Title            string `json:"title"`
	ShortDescription string `json:"short_description"`
	Description      string `json:"description"`
	Perks            string `json:"perks"`
	GoalAmount       int    `json:"goal_amount"`
}
