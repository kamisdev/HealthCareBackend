package controllers

import (
	"context"
	"fmt"
	"net/http"
	"serendipity_backend/SerendipityRequest"
	"serendipity_backend/SerendipityResponse"
	"serendipity_backend/configs"
	"serendipity_backend/models"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var marketplaceCollection *mongo.Collection = configs.GetCollection(configs.DB, "marketplace")

func CreatNewMarketplace() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		email, exists := ctx.Get("email")
		if !exists {
			ctx.JSON(http.StatusBadGateway, gin.H{"error": SerendipityResponse.NOT_LOGGED_IN, "status": "failed"})
			return
		}
		role, exists := ctx.Get("role")
		user_role, _ := strconv.Atoi(fmt.Sprintf("%v", role))
		fmt.Printf("User Role for Update => %v", user_role)

		if user_role == 3 {
			ctx.JSON(http.StatusBadGateway, gin.H{"error": SerendipityResponse.NOT_PERMISSION_ALLOWED, "status": "Permission is not arranged."})
			return
		}
		user_email := fmt.Sprintf("User Email %v", email)
		fmt.Printf("User Email => %v", user_email)
		c, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var input SerendipityRequest.AddNewMarketplace
		defer cancel()

		if err := ctx.ShouldBindJSON(&input); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "status": "Error Occurred in JSON Binding."})
			return
		}

		if validationErr := validate.Struct(&input); validationErr != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error(), "status": "Error Occurred in validation of request."})
			return
		}

		newMarketplaceType := models.MarketPlace{
			ID:               primitive.NewObjectID(),
			Title:            input.Title,
			CoverLetterImage: input.CoverLetterImage,
			Type:             input.Type,
		}

		curMarketplace, err := newMarketplaceType.SaveMarketplace(c)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "status": "Error occurred in creating a new marketplace"})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"status": true, "payload": curMarketplace})
	}
}

func GetAllMarketplaces() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		email, exists := ctx.Get("email")
		if !exists {
			ctx.JSON(http.StatusBadGateway, gin.H{"error": SerendipityResponse.NOT_LOGGED_IN, "status": "failed"})
			return
		}
		c, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		user_email := fmt.Sprint(email)
		fmt.Printf("User Email for Update is %v", user_email)
		defer cancel()

		results, err := models.GetAllMarketplaces(c)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "status": "failed"})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"status": true, "payload": results})
	}
}

func CreateNewMarketplaceItem() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		email, exists := ctx.Get("email")
		if !exists {
			ctx.JSON(http.StatusBadGateway, gin.H{"error": SerendipityResponse.NOT_LOGGED_IN, "status": "failed"})
			return
		}
		role, exists := ctx.Get("role")
		user_role, _ := strconv.Atoi(fmt.Sprintf("%v", role))
		fmt.Printf("User Role for Update => %v", user_role)

		if user_role == 3 {
			ctx.JSON(http.StatusBadGateway, gin.H{"error": SerendipityResponse.NOT_PERMISSION_ALLOWED, "status": "Permission is not arranged."})
			return
		}
		user_email := fmt.Sprintf("User Email %v", email)
		fmt.Printf("User Email => %v", user_email)
		c, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var input SerendipityRequest.AddNewMarketplaceItem
		defer cancel()

		if err := ctx.ShouldBindJSON(&input); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "status": "Error Occurred in JSON Binding."})
			return
		}

		if validationErr := validate.Struct(&input); validationErr != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error(), "status": "Error Occurred in validation of request."})
			return
		}

		newMarketplaceItem := models.MarketplaceItem{
			ID:              primitive.NewObjectID(),
			Title:           input.Title,
			Logo:            input.Logo,
			Description:     input.Description,
			Link:            input.Link,
			MarketplaceType: input.MarketplaceType,
		}

		curMarketplaceItem, err := newMarketplaceItem.SaveMarketplaceItem(c)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "status": "Error occurred in creating a new marketplace"})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"status": true, "payload": curMarketplaceItem})
	}
}

func GetAllMarketplaceItemsWithType() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		email, exists := ctx.Get("email")
		if !exists {
			ctx.JSON(http.StatusBadGateway, gin.H{"error": SerendipityResponse.NOT_LOGGED_IN, "status": "failed"})
			return
		}
		c, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		marketplaceId := ctx.Param("marketplaceId")
		user_email := fmt.Sprint(email)
		fmt.Printf("User Email for Update is %v", user_email)
		defer cancel()

		objMarketplaceId, err := strconv.Atoi(marketplaceId)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "status": "failed"})
			return
		}

		limit := ctx.Request.URL.Query().Get("results")
		page := ctx.Request.URL.Query().Get("page")
		sortField := ctx.Request.URL.Query().Get("sortField")
		sortOrder := ctx.Request.URL.Query().Get("sortOrder")

		convertedLimit, er := strconv.Atoi(limit)
		if er != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": er.Error(), "status": "failed"})
			return
		}
		convertedPage, err := strconv.Atoi(page)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "status": "failed"})
			return
		}
		result, errr := models.GetAllMarketplaceItemsWithType(convertedLimit, convertedPage, sortField, sortOrder, objMarketplaceId, c)

		if errr != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": errr.Error(), "status": "Error occurred in getting marketplace items."})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"status": true, "payload": result})
	}
}
