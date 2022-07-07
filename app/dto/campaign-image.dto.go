package dto

type CreateCampaignImageInput struct {
	Campaign  string `form:"campaign"`
	IsPrimary bool   `form:"is_primary"`
}
