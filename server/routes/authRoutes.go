package routes

import (
	"care-taker/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.RouterGroup) {
	r.GET("/:provider", controllers.BeginAuth)
	r.GET("/:provider/callback", controllers.AuthCallback)
	r.GET("/logout", controllers.Logout)
}
