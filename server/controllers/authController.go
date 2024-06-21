package controllers

import (
	"care-taker/database"
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
)

func BeginAuth(c *gin.Context) {
	provider := c.Param("provider")
	fmt.Println(provider)
	//c.Request = c.Request.WithContext(context.WithValue(context.Background(), "provider", provider))
	ctx:=context.WithValue(c.Request.Context(), "provider", provider)
	c.Request = c.Request.WithContext(ctx)
	gothic.BeginAuthHandler(c.Writer, c.Request)
}

func AuthCallback(c *gin.Context) {
	user, err := gothic.CompleteUserAuth(c.Writer, c.Request)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "authentication failed: " + err.Error()})
		return
	}

	session := sessions.Default(c)
	session.Set("user", user.UserID)
	if err := session.Save(); err != nil {
		log.Println("Failed to save session:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save session: " + err.Error()})
		return
	}

	userEntry := database.UserCollection.FindOne(c.Request.Context(), gin.H{"email": user.Email})
	fmt.Println("userEntry", userEntry)
	if userEntry.Err() != nil {
		_, err := database.UserCollection.InsertOne(c.Request.Context(), gin.H{
			"email":    user.Email,
			"name":     user.Name,
			"googleId": user.UserID,
			"picture":  user.AvatarURL,
			"isAdmin":  false,
		})
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create user: " + err.Error()})
			return
		}
	}

	c.Redirect(http.StatusTemporaryRedirect, "http://localhost:3000")
}

func Logout(c *gin.Context) {
	err := gothic.Logout(c.Writer, c.Request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to logout: " + err.Error()})
		return
	}

	session := sessions.Default(c)
	session.Clear()
	if err := session.Save(); err != nil {
		log.Println("Failed to save session:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save session: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "logged out successfully"})
}
