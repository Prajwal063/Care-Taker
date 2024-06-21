package helpers

import "github.com/gin-gonic/gin"

func SliceContains(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

func SendResponse(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, gin.H{"error": message})
}