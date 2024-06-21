package controllers

import (
	"net/http"

	"care-taker/database"
	"care-taker/helpers"
	"care-taker/models"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	userId := c.MustGet("userId").(string)

	var user models.User
	err := database.UserCollection.FindOne(c.Request.Context(), gin.H{"googleId": userId}).Decode(&user)
	if err != nil {
		helpers.SendResponse(c, http.StatusInternalServerError, "Failed to fetch user: "+err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user.ToJSON()})
}

func UpdateUser(c *gin.Context) {
	userId := c.MustGet("userId").(string)

	var user models.User
	err := c.BindJSON(&user)
	if err != nil {
		helpers.SendResponse(c, http.StatusBadRequest, "Invalid request payload: "+err.Error())
		return
	}

	_, err = database.UserCollection.UpdateOne(c.Request.Context(), gin.H{"googleId": userId}, gin.H{"$set": user})
	if err != nil {
		helpers.SendResponse(c, http.StatusInternalServerError, "Failed to update user: "+err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated", "user": user.ToJSON()})
}
