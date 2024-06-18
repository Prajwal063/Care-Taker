package controllers

import (
	"context"
	"net/http"

	"care-taker/database"
	"care-taker/models"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	userId := c.MustGet("userId").(string)

	var user models.User
	err := database.UserCollection.FindOne(context.TODO(), gin.H{"googleId": userId}).Decode(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user.ToJSON()})
}

func UpdateUser(c *gin.Context) {
	userId := c.MustGet("userId").(string)

	var user models.User
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request: " + err.Error()})
		return
	}

	_, err = database.UserCollection.UpdateOne(context.TODO(), gin.H{"googleId": userId}, gin.H{"$set": user})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated", "user": user.ToJSON()})
}
