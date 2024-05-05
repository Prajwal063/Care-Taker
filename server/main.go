package main

import (
	"fmt"
	"log"
	"os"
	"server/database"
	"server/initializers"

	helmet "github.com/danielkov/gin-helmet"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.Env()
	database.Client = initializers.Database()
}

func main() {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(gin.ErrorLogger())
	router.Use(helmet.Default())

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	err := router.Run(":" + port)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("Server is running on port: ", port)
}
