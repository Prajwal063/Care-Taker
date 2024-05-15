package main

import (
	"fmt"
	"log"
	"os"

	"care-taker/database"
	"care-taker/initializers"
	"care-taker/routes"

	helmet "github.com/danielkov/gin-helmet"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.Env()
	database.Client = initializers.Database()
	initializers.Collections()
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

	routes.ServiceRoutes(router.Group("/service"))
	routes.EventRoutes(router.Group("/event"))

	err := router.Run(":" + port)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("Server is running on port: ", port)
}
