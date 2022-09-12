package models

import (
	"context"
	"errors"
	"serendipity_backend/SerendipityResponse"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MarketPlace struct {
	ID               primitive.ObjectID `json:"id,omitempty"`
	Title            string             `json:"title,omitempty"`
	CoverLetterImage string             `json:"coverLetterImage,omitempty"`
	Type             int                `json:"type,omitempty"` //-> 1: Food, 2: Assistive Devices 3: Services & Experiences 4: Coaching 5: Foundations
}

type MarketplaceItem struct {
	ID              primitive.ObjectID `json:"id,omitempty"`
	Title           string             `json:"title,omitempty"`
	Logo            string             `json:"logo,omitempty"`
	Link            string             `json:"link,omitempty"`
	Description     string             `json:"description,omitempty"`
	MarketplaceType int                `json:"marketplaceType,omitempty"`
}

func (newMarketplace MarketPlace) SaveMarketplace(c context.Context) (MarketPlace, error) {
	marketPlace, err := marketplaceCollection.InsertOne(c, newMarketplace)
	if err != nil {
		if er, ok := err.(mongo.WriteException); ok && er.WriteErrors[0].Code == 11000 {
			return MarketPlace{}, errors.New("Marketplace Already Exists")
		}
	}

	opt := options.Index()
	opt.SetUnique(true)

	index := mongo.IndexModel{Keys: bson.M{"type": 1}, Options: opt}

	if _, err := marketplaceCollection.Indexes().CreateOne(c, index); err != nil {
		return MarketPlace{}, errors.New("could not create index of Marketplace Type")
	}

	var curMarketplace MarketPlace
	marketplaceCollection.FindOne(c, bson.M{"_id": marketPlace.InsertedID}).Decode(&curMarketplace)
	return curMarketplace, nil
}

func GetAllMarketplaces(c context.Context) ([]MarketPlace, error) {
	var marketplaces []MarketPlace
	results, err := marketplaceCollection.Find(c, bson.M{})
	if err != nil {
		return []MarketPlace{}, errors.New(SerendipityResponse.ERROR_GET_MARKETPLACES)
	}
	defer results.Close(c)

	for results.Next(c) {
		var singleMarketplace MarketPlace
		if err := results.Decode(&singleMarketplace); err != nil {
			return []MarketPlace{}, errors.New(SerendipityResponse.ERROR_DECODE_MARKETPLACE)
		}
		marketplaces = append(marketplaces, singleMarketplace)
	}
	return marketplaces, nil
}

func (newMarketplaceItem MarketplaceItem) SaveMarketplaceItem(c context.Context) (MarketplaceItem, error) {
	marketplaceItem, err := marketplaceItemCollection.InsertOne(c, newMarketplaceItem)
	if err != nil {
		if er, ok := err.(mongo.WriteException); ok && er.WriteErrors[0].Code == 11000 {
			return MarketplaceItem{}, errors.New("Marketplace Item Already Exists")
		}
	}

	opt := options.Index()
	opt.SetUnique(true)

	index := mongo.IndexModel{Keys: bson.M{"title": 1}, Options: opt}

	if _, err := marketplaceItemCollection.Indexes().CreateOne(c, index); err != nil {
		return MarketplaceItem{}, errors.New("could not create index of Marketplace Item title")
	}

	var curMarketplaceItem MarketplaceItem
	marketplaceItemCollection.FindOne(c, bson.M{"_id": marketplaceItem.InsertedID}).Decode(&curMarketplaceItem)
	return curMarketplaceItem, nil
}

func GetAllMarketplaceItemsWithType(limit int, page int, sortField string, sortOrder string, objId int, c context.Context) ([]MarketplaceItem, error) {

	skipNumber := (page - 1) * limit
	pageOptions := options.Find()
	if sortOrder == "ascend" {
		pageOptions.SetSort(bson.D{primitive.E{Key: sortField, Value: 1}})
	} else if sortOrder == "descend" {
		pageOptions.SetSort(bson.D{primitive.E{Key: sortField, Value: -1}})
	}
	pageOptions.SetSkip(int64(skipNumber))
	pageOptions.SetLimit(int64(limit))

	var marketplaceItems []MarketplaceItem
	result, err := marketplaceItemCollection.Find(c, bson.M{"marketplacetype": objId}, pageOptions)
	if err != nil {
		return []MarketplaceItem{}, errors.New(SerendipityResponse.ERROR_GETALL_MARKETPLACEITEMS)
	}

	defer result.Close(c)

	for result.Next(c) {
		var singleMarketplaceItem MarketplaceItem
		if err := result.Decode(&singleMarketplaceItem); err != nil {
			return []MarketplaceItem{}, errors.New(SerendipityResponse.ERROR_DECODE_MARKETPLACE_ITEM)
		}
		marketplaceItems = append(marketplaceItems, singleMarketplaceItem)
	}
	return marketplaceItems, nil

}
