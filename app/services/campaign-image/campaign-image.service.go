package campaign_image

import (
	"go-gin-mongodb-clean-architecture/app/dto"
	"go-gin-mongodb-clean-architecture/app/entities"
	campaignImageRepo "go-gin-mongodb-clean-architecture/app/repositories/campaign-image"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Service interface {
	CreateCampaignImage(input dto.CreateCampaignImageInput, fileLocation string) (string, error)
}

type service struct {
	campaignImageRepository campaignImageRepo.Repository
}

func NewService(campaignImageRepository campaignImageRepo.Repository) *service {
	return &service{campaignImageRepository: campaignImageRepository}
}

func (s *service) CreateCampaignImage(input dto.CreateCampaignImageInput, fileLocation string) (string, error) {
	campaignImage := entities.CampaignImage{}
	objID, _ := primitive.ObjectIDFromHex(input.Campaign)
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
