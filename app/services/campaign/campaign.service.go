package campaign

import (
	"fmt"
	"github.com/gosimple/slug"
	"go-gin-mongodb-clean-architecture/app/dto"
	"go-gin-mongodb-clean-architecture/app/entities"
	"go-gin-mongodb-clean-architecture/app/repositories/campaign"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type CampaignService interface {
	CreateCampaign(input dto.CreateCampaignInput) (string, error)
}

type campaignService struct {
	campaignRepository campaign.Repository
}

func NewService(campaignRepository campaign.Repository) *campaignService {
	return &campaignService{
		campaignRepository: campaignRepository,
	}
}

func (s *campaignService) Create(input dto.CreateCampaignInput) (string, error) {
	campaign := entities.Campaign{}
	campaign.ID = primitive.NewObjectID()
	campaign.Title = input.Title
	campaign.ShortDescription = input.ShortDescription
	campaign.Description = input.Description
	campaign.Perks = input.Perks
	campaign.GoalAmount = input.GoalAmount
	campaign.User = input.User
	campaign.GoalAmount = input.GoalAmount
	campaign.CurrentAmount = 0

	campaignID, _ := primitive.ObjectIDFromHex(campaign.ID.Hex())
	campaignSlug := slug.Make(fmt.Sprintf("%s-%s", input.Title, campaignID))
	campaign.Slug = campaignSlug

	campaign.CreatedAt = time.Now()
	campaign.UpdatedAt = time.Now()

	newCampaign, err := s.campaignRepository.Create(campaign)
	if err != nil {
		return "", err
	}

	return newCampaign, nil
}
