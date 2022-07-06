package campaign

import (
	"context"
	"fmt"
	"go-gin-mongodb-clean-architecture/app/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type Repository interface {
	Create(campaign entities.Campaign) (string, error)
	FindBySlug(slug string) (entities.Campaign, error)
	FindByUser(User primitive.ObjectID) ([]entities.Campaign, error)
	FindAll() ([]entities.Campaign, error)
	FindByID(ID primitive.ObjectID) (entities.Campaign, error)
	UpdateByID(ID primitive.ObjectID, campaign entities.Campaign) (int, error)
	UpdateBySlug(slug string, campaign entities.Campaign) (int, error)
}

type repository struct {
	campaignCollection *mongo.Collection
}

func NewRepository(campaignCollection *mongo.Collection) *repository {
	return &repository{campaignCollection: campaignCollection}
}

func (r *repository) Create(campaign entities.Campaign) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := r.campaignCollection.InsertOne(ctx, campaign)
	if err != nil {
		return "", err
	}

	insertedID := fmt.Sprintf("%s", result.InsertedID)
	return insertedID, nil
}

func (r *repository) FindBySlug(slug string) (entities.Campaign, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var campaign entities.Campaign
	filter := bson.M{"slug": slug}
	err := r.campaignCollection.FindOne(ctx, filter).Decode(&campaign)
	if err != nil {
		return campaign, err
	}

	return campaign, nil
}

func (r *repository) FindByUser(User primitive.ObjectID) ([]entities.Campaign, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var campaigns []entities.Campaign
	filter := bson.M{"user": User}
	aggSearch := bson.M{"$match": filter}
	lookupStage := bson.M{"$lookup": bson.M{
		"from":         "users",
		"localField":   "user",
		"foreignField": "_id",
		"as":           "users",
	}}
	result, err := r.campaignCollection.Aggregate(ctx, []bson.M{aggSearch, lookupStage})
	if err != nil {
		return campaigns, err
	}

	for result.Next(ctx) {
		var campaign entities.Campaign
		err := result.Decode(&campaign)
		if err != nil {
			return campaigns, err
		}
		campaigns = append(campaigns, campaign)
	}
	return campaigns, err
}

func (r *repository) FindAll() ([]entities.Campaign, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var campaigns []entities.Campaign
	filter := bson.M{}
	aggSearch := bson.M{"$match": filter}
	aggPopulate := bson.M{"$lookup": bson.M{
		"from":         "users",
		"localField":   "user",
		"foreignField": "_id",
		"as":           "users",
	}}
	cursor, err := r.campaignCollection.Aggregate(ctx, []bson.M{aggSearch, aggPopulate})
	if err != nil {
		return campaigns, err
	}

	for cursor.Next(ctx) {
		var campaign entities.Campaign
		err = cursor.Decode(&campaign)
		if err != nil {
			return campaigns, err
		}
		campaigns = append(campaigns, campaign)
	}

	return campaigns, nil
}

func (r *repository) FindByID(ID primitive.ObjectID) (entities.Campaign, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var campaign entities.Campaign

	filter := bson.M{"_id": ID}
	aggSearch := bson.M{"$match": filter}
	aggPopulate := bson.M{"$lookup": bson.M{
		"from":         "users",
		"localField":   "user",
		"foreignField": "_id",
		"as":           "users",
	}}
	cursor, err := r.campaignCollection.Aggregate(ctx, []bson.M{aggSearch, aggPopulate})
	if err != nil {
		return campaign, err
	}

	for cursor.Next(ctx) {
		err := cursor.Decode(&campaign)
		if err != nil {
			return campaign, err
		}
		return campaign, nil
	}
	return campaign, nil
}

func (r *repository) UpdateByID(ID primitive.ObjectID, campaign entities.Campaign) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"_id": ID}
	update := bson.M{
		"$set": bson.M{
			"title":             campaign.Title,
			"short_description": campaign.ShortDescription,
			"description":       campaign.Description,
			"perks":             campaign.Perks,
			"backer_count":      campaign.BackerCount,
			"goal_amount":       campaign.GoalAmount,
			"current_amount":    campaign.CurrentAmount,
			"slug":              campaign.Slug,
			"updated_at":        campaign.UpdatedAt,
			"created_at":        campaign.CreatedAt,
		},
	}
	updateResult, err := r.campaignCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return 0, err
	}

	return int(updateResult.ModifiedCount), nil
}

func (r *repository) UpdateBySlug(slug string, campaign entities.Campaign) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"slug": slug}
	update := bson.M{
		"$set": bson.M{
			"title":             campaign.Title,
			"short_description": campaign.ShortDescription,
			"description":       campaign.Description,
			"perks":             campaign.Perks,
			"backer_count":      campaign.BackerCount,
			"goal_amount":       campaign.GoalAmount,
			"current_amount":    campaign.CurrentAmount,
			"slug":              campaign.Slug,
			"updated_at":        campaign.UpdatedAt,
			"created_at":        campaign.CreatedAt,
		},
	}
	updateResult, err := r.campaignCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return 0, err
	}

	return int(updateResult.ModifiedCount), nil
}
