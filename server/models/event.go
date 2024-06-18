package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var EventStatus = []string{"UPCOMING", "ONGOING", "COMPLETED", "CANCELLED"}

// Event represents the MongoDB schema for the Event collection.
type Event struct {
	ID              primitive.ObjectID `bson:"_id,omitempty"`
	Title           string             `bson:"title,omitempty" validate:"required,min=3,max=32"`
	Description     string             `bson:"description,omitempty" validate:"required,min=3,max=500"`
	Email           string             `bson:"email,omitempty" validate:"required,email"`
	Phone           string             `bson:"phone,omitempty" validate:"required,len=10"`
	Address         string             `bson:"address,omitempty" validate:"required,max=100"`
	Picture         string             `bson:"picture,omitempty" validate:"required"`
	Date            string             `bson:"date,omitempty" validate:"required"`
	Time            string             `bson:"time,omitempty" validate:"required"`
	Price           float64            `bson:"price,omitempty" validate:"required"`
	Status          string             `bson:"status,omitempty" validate:"required,oneof=UPCOMING ONGOING COMPLETED CANCELLED"`
	RegisteredUsers []string           `bson:"registeredUsers,omitempty"`
	CreatedAt       time.Time          `bson:"createdAt,omitempty"`
	UpdatedAt       time.Time          `bson:"updatedAt,omitempty"`
}

// EventResponse represents the response format for Event.
type EventResponse struct {
	ID              string    `json:"id,omitempty"`
	Title           string    `json:"title,omitempty"`
	Description     string    `json:"description,omitempty"`
	Email           string    `json:"email,omitempty"`
	Phone           string    `json:"phone,omitempty"`
	Address         string    `json:"address,omitempty"`
	Picture         string    `json:"picture,omitempty"`
	Date            string    `json:"date,omitempty"`
	Time            string    `json:"time,omitempty"`
	Price           float64   `json:"price,omitempty"`
	Status          string    `json:"status,omitempty"`
	RegisteredUsers []string  `json:"registeredUsers,omitempty"`
	CreatedAt       time.Time `json:"createdAt,omitempty"`
	UpdatedAt       time.Time `json:"updatedAt,omitempty"`
}

// ToEventResponse converts an Event to an EventResponse.
func ToEventResponse(event Event) EventResponse {
	return EventResponse{
		ID:              event.ID.Hex(),
		Title:           event.Title,
		Description:     event.Description,
		Email:           event.Email,
		Phone:           event.Phone,
		Address:         event.Address,
		Picture:         event.Picture,
		Date:            event.Date,
		Time:            event.Time,
		Price:           event.Price,
		Status:          event.Status,
		RegisteredUsers: event.RegisteredUsers,
		CreatedAt:       event.CreatedAt,
		UpdatedAt:       event.UpdatedAt,
	}
}

// ToEventResponseArray converts a slice of Event to a slice of EventResponse.
func ToEventResponseArray(events []Event) []EventResponse {
	var eventResponses []EventResponse
	for _, event := range events {
		eventResponses = append(eventResponses, ToEventResponse(event))
	}
	return eventResponses
}
