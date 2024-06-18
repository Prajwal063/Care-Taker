package controllers

import (
	"context"
	"net/http"

	"care-taker/database"
	"care-taker/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAllServices(c *gin.Context) {
	search := c.Query("search")
	filter := bson.M{}
	if search != "" {
		filter = bson.M{"title": bson.M{"$regex": search, "$options": "i"}}
	}

	cursor, err := database.ServiceCollection.Find(context.TODO(), filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch services"})
		return
	}
	defer cursor.Close(context.TODO())

	var services []models.Service
	if err := cursor.All(context.TODO(), &services); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode services"})
		return
	}
	if len(services) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No services found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Fetched services Successfully!", "data": models.ToServiceResponseArray(services)})
}

func CreateService(c *gin.Context) {
	var service models.Service
	if err := c.BindJSON(&service); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	serviceID, err := database.ServiceCollection.InsertOne(context.TODO(), service)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create service"})
		return
	}
	service.ID = serviceID.InsertedID.(primitive.ObjectID)

	c.JSON(http.StatusCreated, gin.H{"message": "Service created Successfully!", "data": models.ToServiceResponse(service)})
}

func GetServiceById(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid service ID"})
		return
	}
	var service models.Service
	err = database.ServiceCollection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&service)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Service not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Fetched service Successfully!", "data": models.ToServiceResponse(service)})
}

func UpdateService(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid service ID"})
		return
	}

	var updatedService models.Service
	if err := c.BindJSON(&updatedService); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	res, err := database.ServiceCollection.UpdateOne(context.TODO(), bson.M{"_id": id}, bson.M{"$set": updatedService})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update service"})
		return
	}
	if res.MatchedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Service not found"})
		return
	}

	updatedService.ID = id

	c.JSON(http.StatusOK, gin.H{"message": "Service updated Successfully!", "data": models.ToServiceResponse(updatedService)})
}

func DeleteService(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid service ID"})
		return
	}

	res, err := database.ServiceCollection.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete service"})
		return
	}
	if res.DeletedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Service not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Service deleted Successfully!"})
}
