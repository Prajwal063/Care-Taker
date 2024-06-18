package routes

import (
	"care-taker/controllers"
	"care-taker/middlewares"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.RouterGroup) {
	r.GET("/me", middlewares.IsAuthenticated, controllers.GetUser)
	r.PUT("/me", middlewares.IsAuthenticated, controllers.UpdateUser)
}
