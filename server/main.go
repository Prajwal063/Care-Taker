package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"care-taker/database"
	"care-taker/initializers"
	"care-taker/routes"

	helmet "github.com/danielkov/gin-helmet"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.Env()
	database.Client = initializers.Database()
	initializers.Collections()
}

func main() {
	router := gin.New()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // Adjust as needed
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		AllowCredentials: true,
		MaxAge:           24 * time.Hour,
	}))

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(gin.ErrorLogger())
	router.Use(helmet.Default())
	router.Use(sessions.Sessions("session", initializers.AuthStore()))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	routes.AuthRoutes(router.Group("/auth"))
	routes.ServiceRoutes(router.Group("/service"))
	routes.EventRoutes(router.Group("/event"))
	routes.UserRoutes(router.Group("/user"))

	err := router.Run("localhost:" + port)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("Server is running on port: ", port)
}
