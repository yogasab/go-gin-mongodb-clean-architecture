package campaign

import (
	"errors"
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
	GetCampaignsUser(UserID primitive.ObjectID) ([]entities.Campaign, error)
	GetCampaigns(UserID string) ([]entities.Campaign, error)
}

type campaignService struct {
	campaignRepository campaign.Repository
}

func NewService(campaignRepository campaign.Repository) *campaignService {
	return &campaignService{
		campaignRepository: campaignRepository,
	}
}

func (s *campaignService) CreateCampaign(input dto.CreateCampaignInput) (string, error) {
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
	campaignSlug := slug.Make(fmt.Sprintf("%s-%s", input.Title, input.User.Hex()))
	campaign.Slug = campaignSlug
	campaign.CreatedAt = time.Now()
	campaign.UpdatedAt = time.Now()

	isExist, _ := s.campaignRepository.FindBySlug(campaignSlug)
	if isExist.Title != "" {
		return "", errors.New("Campaign is already created, please make another unique campaign")
	}

	newCampaign, err := s.campaignRepository.Create(campaign)
	if err != nil {
		return "", err
	}

	return newCampaign, nil
}

func (s *campaignService) GetCampaignsUser(UserID primitive.ObjectID) ([]entities.Campaign, error) {
	campaignsUser, err := s.campaignRepository.FindByUser(UserID)
	if err != nil {
		return campaignsUser, err
	}

	return campaignsUser, nil
}

func (s *campaignService) GetCampaigns(UserID string) ([]entities.Campaign, error) {
	if UserID == "" {
		campaigns, err := s.campaignRepository.FindAll()
		if err != nil {
			return campaigns, err
		}
		return campaigns, nil
	}

	objID, _ := primitive.ObjectIDFromHex(UserID)
	campaigns, err := s.campaignRepository.FindByUser(objID)
	if len(campaigns) == 0 {
		return campaigns, errors.New("Campaign with correspond user is not found")
	}
	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}
