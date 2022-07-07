package campaign_image

import (
	"errors"
	"go-gin-mongodb-clean-architecture/app/dto"
	"go-gin-mongodb-clean-architecture/app/entities"
	campaignRepo "go-gin-mongodb-clean-architecture/app/repositories/campaign"
	campaignImageRepo "go-gin-mongodb-clean-architecture/app/repositories/campaign-image"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"time"
)

type Service interface {
	CreateCampaignImage(input dto.CreateCampaignImageInput, fileLocation string) (string, error)
}

type service struct {
	campaignImageRepository campaignImageRepo.Repository
	campaignRepository      campaignRepo.Repository
}

func NewService(campaignImageRepository campaignImageRepo.Repository, campaignRepository campaignRepo.Repository) *service {
	return &service{campaignImageRepository: campaignImageRepository, campaignRepository: campaignRepository}
}

func (s *service) CreateCampaignImage(input dto.CreateCampaignImageInput, fileLocation string) (string, error) {
	objID, _ := primitive.ObjectIDFromHex(input.Campaign)
	campaign, err := s.campaignRepository.FindByID(objID)
	if campaign.Title == "" {
		return "", errors.New("Campaign is not found")
	}
	if err != nil {
		return "", err
	}
	log.Println(campaign.User)
	log.Println(input.User.ID)
	if campaign.User != input.User.ID {
		return "", errors.New("You are not the owner of the campaign")
	}

	campaignImage := entities.CampaignImage{}
	campaignImage.ID = input.ID
	campaignImage.Campaign = objID
	campaignImage.Filename = fileLocation
	campaignImage.IsPrimary = input.IsPrimary
	campaignImage.CreatedAt = time.Now()
	campaignImage.UpdatedAt = time.Now()

	newCampaignImage, err := s.campaignImageRepository.Create(campaignImage)
	if err != nil {
		return "", err
	}

	return newCampaignImage, nil
}
