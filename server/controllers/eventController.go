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

func GetAllEvents(c *gin.Context) {
	search := c.Query("search")
	filter := bson.M{}
	if search != "" {
		filter = bson.M{"title": bson.M{"$regex": search}}
	}

	cursor, err := database.EventCollection.Find(context.TODO(), filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch events"})
		return
	}
	defer cursor.Close(context.TODO())

	var events []models.Event
	if err := cursor.All(context.TODO(), &events); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode events"})
		return
	}
	if len(events) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No events found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Fetched events Successfully!", "data": models.ToEventResponseArray(events)})
}

func CreateEvent(c *gin.Context) {
	var event models.Event
	if err := c.BindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	eventID, err := database.EventCollection.InsertOne(context.TODO(), event)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create event"})
		return
	}
	event.ID = eventID.InsertedID.(primitive.ObjectID)

	c.JSON(http.StatusCreated, gin.H{"message": "Event created Successfully!", "data": models.ToEventResponse(event)})
}

func GetEventById(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}
	var event models.Event
	err = database.EventCollection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&event)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Fetched event Successfully!", "data": models.ToEventResponse(event)})
}

func UpdateEvent(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}

	var updatedEvent models.Event
	if err := c.BindJSON(&updatedEvent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	res, err := database.EventCollection.UpdateOne(context.TODO(), bson.M{"_id": id}, bson.M{"$set": updatedEvent})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update event"})
		return
	}
	if res.MatchedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}

	updatedEvent.ID = id

	c.JSON(http.StatusOK, gin.H{"message": "Event updated Successfully!", "data": models.ToEventResponse(updatedEvent)})
}

func DeleteEvent(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}

	res, err := database.EventCollection.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete event"})
		return
	}
	if res.DeletedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Event deleted Successfully!"})
}
