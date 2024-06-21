package controllers

import (
	"net/http"

	"care-taker/database"
	"care-taker/helpers"
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

	cursor, err := database.ServiceCollection.Find(c.Request.Context(), filter)
	if err != nil {
		helpers.SendResponse(c, http.StatusInternalServerError, "Failed to fetch services")
		return
	}
	defer cursor.Close(c.Request.Context())

	var services []models.Service
	if err := cursor.All(c.Request.Context(), &services); err != nil {
		helpers.SendResponse(c, http.StatusInternalServerError, "Failed to decode services")
		return
	}
	if len(services) == 0 {
		helpers.SendResponse(c, http.StatusNotFound, "No services found")
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Fetched services Successfully!", "data": models.ToServiceResponseArray(services)})
}

func CreateService(c *gin.Context) {
	var service models.Service
	if err := c.BindJSON(&service); err != nil {
		helpers.SendResponse(c, http.StatusBadRequest, "Invalid request payload")
		return
	}

	serviceID, err := database.ServiceCollection.InsertOne(c.Request.Context(), service)
	if err != nil {
		helpers.SendResponse(c, http.StatusInternalServerError, "Failed to create service")
		return
	}
	service.ID = serviceID.InsertedID.(primitive.ObjectID)

	c.JSON(http.StatusCreated, gin.H{"message": "Service created Successfully!", "data": models.ToServiceResponse(service)})
}

func GetServiceById(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		helpers.SendResponse(c, http.StatusBadRequest, "Invalid service ID")
		return
	}
	var service models.Service
	err = database.ServiceCollection.FindOne(c.Request.Context(), bson.M{"_id": id}).Decode(&service)
	if err != nil {
		helpers.SendResponse(c, http.StatusNotFound, "Service not found")
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Fetched service Successfully!", "data": models.ToServiceResponse(service)})
}

func UpdateService(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		helpers.SendResponse(c, http.StatusBadRequest, "Invalid service ID")
		return
	}

	var updatedService models.Service
	if err := c.BindJSON(&updatedService); err != nil {
		helpers.SendResponse(c, http.StatusBadRequest, "Invalid request payload")
		return
	}

	res, err := database.ServiceCollection.UpdateOne(c.Request.Context(), bson.M{"_id": id}, bson.M{"$set": updatedService})
	if err != nil {
		helpers.SendResponse(c, http.StatusInternalServerError, "Failed to update service")
		return
	}
	if res.MatchedCount == 0 {
		helpers.SendResponse(c, http.StatusNotFound, "Service not found")
		return
	}

	updatedService.ID = id

	c.JSON(http.StatusOK, gin.H{"message": "Service updated Successfully!", "data": models.ToServiceResponse(updatedService)})
}

func DeleteService(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		helpers.SendResponse(c, http.StatusBadRequest, "Invalid service ID")
		return
	}

	res, err := database.ServiceCollection.DeleteOne(c.Request.Context(), bson.M{"_id": id})
	if err != nil {
		helpers.SendResponse(c, http.StatusInternalServerError, "Failed to delete service")
		return
	}
	if res.DeletedCount == 0 {
		helpers.SendResponse(c, http.StatusNotFound, "Service not found")
		return
	}

	helpers.SendResponse(c, http.StatusOK, "Service deleted Successfully!")
}
