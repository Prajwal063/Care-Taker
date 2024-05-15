package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Service represents the MongoDB schema for the Service collection.
type Service struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Title       string             `bson:"title,omitempty" validate:"required,min=3,max=32"`
	Description string             `bson:"description,omitempty" validate:"required,min=3,max=500"`
	Email       string             `bson:"email,omitempty" validate:"required,email"`
	Phone       string             `bson:"phone,omitempty" validate:"required,len=10"`
	Address     string             `bson:"address,omitempty" validate:"required,max=100"`
	Picture     string             `bson:"picture,omitempty" validate:"required"`
	CreatedAt   time.Time          `bson:"createdAt,omitempty"`
	UpdatedAt   time.Time          `bson:"updatedAt,omitempty"`
}

// ServiceResponse represents the response format for Service.
type ServiceResponse struct {
	ID          string    `json:"id,omitempty"`
	Title       string    `json:"title,omitempty"`
	Description string    `json:"description,omitempty"`
	Email       string    `json:"email,omitempty"`
	Phone       string    `json:"phone,omitempty"`
	Address     string    `json:"address,omitempty"`
	Picture     string    `json:"picture,omitempty"`
	CreatedAt   time.Time `json:"createdAt,omitempty"`
	UpdatedAt   time.Time `json:"updatedAt,omitempty"`
}

// ToServiceResponse converts an Service to an ServiceResponse.
func ToServiceResponse(service Service) ServiceResponse {
	return ServiceResponse{
		ID:          service.ID.Hex(),
		Title:       service.Title,
		Description: service.Description,
		Email:       service.Email,
		Phone:       service.Phone,
		Address:     service.Address,
		Picture:     service.Picture,
		CreatedAt:   service.CreatedAt,
		UpdatedAt:   service.UpdatedAt,
	}
}

// ToServiceResponseArray converts a slice of Service to a slice of ServiceResponse.
func ToServiceResponseArray(services []Service) []ServiceResponse {
	var serviceResponses []ServiceResponse
	for _, service := range services {
		serviceResponses = append(serviceResponses, ToServiceResponse(service))
	}
	return serviceResponses
}
