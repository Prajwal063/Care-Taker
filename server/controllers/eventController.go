package controllers

import (
	"fmt"
	"net/http"

	"care-taker/database"
	"care-taker/helpers"
	"care-taker/models"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Helper function to check if a user is logged in and fetch user details
func getSessionUser(c *gin.Context) (*models.User, error) {
	session := sessions.Default(c)
	userID := session.Get("user")
	if userID == nil {
		return nil, fmt.Errorf("user not logged in")
	}

	var user models.User
	err := database.UserCollection.FindOne(c.Request.Context(), bson.M{"googleId": userID}).Decode(&user)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch user details: %v", err)
	}

	return &user, nil
}

func GetAllEvents(c *gin.Context) {
	search := c.Query("search")
	filter := bson.M{}
	if search != "" {
		filter = bson.M{"title": bson.M{"$regex": search, "$options": "i"}}
	}

	cursor, err := database.EventCollection.Find(c.Request.Context(), filter)
	if err != nil {
		helpers.SendResponse(c, http.StatusInternalServerError, "Failed to fetch events")
		return
	}
	defer cursor.Close(c.Request.Context())

	var events []models.Event
	if err := cursor.All(c.Request.Context(), &events); err != nil {
		helpers.SendResponse(c, http.StatusInternalServerError, "Failed to decode events")
		return
	}
	if len(events) == 0 {
		helpers.SendResponse(c, http.StatusNotFound, "No events found")
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Fetched events Successfully!", "data": models.ToEventResponseArray(events)})
}

func CreateEvent(c *gin.Context) {
	var event models.Event
	if err := c.BindJSON(&event); err != nil {
		helpers.SendResponse(c, http.StatusBadRequest, "Invalid request payload")
		return
	}

	eventID, err := database.EventCollection.InsertOne(c.Request.Context(), event)
	if err != nil {
		helpers.SendResponse(c, http.StatusInternalServerError, "Failed to create event")
		return
	}
	event.ID = eventID.InsertedID.(primitive.ObjectID)

	c.JSON(http.StatusCreated, gin.H{"message": "Event created Successfully!", "data": models.ToEventResponse(event)})
}

func GetEventById(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		helpers.SendResponse(c, http.StatusBadRequest, "Invalid event ID")
		return
	}
	var event models.Event
	err = database.EventCollection.FindOne(c.Request.Context(), bson.M{"_id": id}).Decode(&event)
	if err != nil {
		helpers.SendResponse(c, http.StatusNotFound, "Event not found")
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Fetched event Successfully!", "data": models.ToEventResponse(event)})
}

func UpdateEvent(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		helpers.SendResponse(c, http.StatusBadRequest, "Invalid event ID")
		return
	}

	var updatedEvent models.Event
	if err := c.BindJSON(&updatedEvent); err != nil {
		helpers.SendResponse(c, http.StatusBadRequest, "Invalid request payload")
		return
	}

	res, err := database.EventCollection.UpdateOne(c.Request.Context(), bson.M{"_id": id}, bson.M{"$set": updatedEvent})
	if err != nil {
		helpers.SendResponse(c, http.StatusInternalServerError, "Failed to update event")
		return
	}
	if res.MatchedCount == 0 {
		helpers.SendResponse(c, http.StatusNotFound, "Event not found")
		return
	}

	updatedEvent.ID = id

	c.JSON(http.StatusOK, gin.H{"message": "Event updated Successfully!", "data": models.ToEventResponse(updatedEvent)})
}

func DeleteEvent(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		helpers.SendResponse(c, http.StatusBadRequest, "Invalid event ID")
		return
	}

	res, err := database.EventCollection.DeleteOne(c.Request.Context(), bson.M{"_id": id})
	if err != nil {
		helpers.SendResponse(c, http.StatusInternalServerError, "Failed to delete event")
		return
	}
	if res.DeletedCount == 0 {
		helpers.SendResponse(c, http.StatusNotFound, "Event not found")
		return
	}

	helpers.SendResponse(c, http.StatusOK, "Event deleted Successfully!")
}

func RegisterToEvent(c *gin.Context) {
	eventID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		helpers.SendResponse(c, http.StatusBadRequest, "Invalid event ID")
		return
	}
	user, err := getSessionUser(c)
	if err != nil {
		helpers.SendResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	var event models.Event
	err = database.EventCollection.FindOne(c.Request.Context(), bson.M{"_id": eventID}).Decode(&event)
	if err != nil {
		helpers.SendResponse(c, http.StatusInternalServerError, "Failed to fetch event details")
		return
	}

	if event.Status == "closed" {
		helpers.SendResponse(c, http.StatusBadRequest, "Event registration closed")
		return
	}

	if helpers.SliceContains(event.RegisteredUsers, user.ID.Hex()) {
		helpers.SendResponse(c, http.StatusBadRequest, "USer already registered for event")
		return
	}

	res, err := database.EventCollection.UpdateOne(c.Request.Context(), bson.M{"_id": eventID}, bson.M{"$addToSet": bson.M{"participants": user.ID}})
	if err != nil {
		helpers.SendResponse(c, http.StatusInternalServerError, "Failed to register for event")
		return
	}
	if res.MatchedCount == 0 {
		helpers.SendResponse(c, http.StatusNotFound, "Event not found")
		return
	}

	helpers.SendResponse(c, http.StatusOK, "Registered for event Successfully!")
}

func UnregisterFromEvent(c *gin.Context) {
	eventID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		helpers.SendResponse(c, http.StatusBadRequest, "Invalid event ID")
		return
	}

	user, err := getSessionUser(c)
	if err != nil {
		helpers.SendResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	res, err := database.EventCollection.UpdateOne(c.Request.Context(), bson.M{"_id": eventID}, bson.M{"$pull": bson.M{"participants": user.ID}})
	if err != nil {
		helpers.SendResponse(c, http.StatusInternalServerError, "Failed to unregister from event")
		return
	}
	if res.MatchedCount == 0 {
		helpers.SendResponse(c, http.StatusNotFound, "Event not found")
		return
	}

	helpers.SendResponse(c, http.StatusOK, "Unregistered from event Successfully!")
}
