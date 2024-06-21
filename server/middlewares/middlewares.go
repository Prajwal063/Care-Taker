package middlewares

import (
	"net/http"

	"care-taker/database"
	"care-taker/models"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func IsAuthenticated(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get("user")
	if user == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		c.Abort()
		return
	}
	c.Set("userId", user)
	c.Next()
}

func IsAdmin(c *gin.Context) {
	session := sessions.Default(c)
	userId := session.Get("user")
	if userId == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		c.Abort()
		return
	}

	var user models.User
	err := database.UserCollection.FindOne(c.Request.Context(), gin.H{"googleId": userId}).Decode(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to verify user: " + err.Error()})
		c.Abort()
		return
	}

	if user.IsAdmin {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		c.Abort()
		return
	}
	c.Set("userId", userId)
	c.Next()
}
