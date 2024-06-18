package routes

import (
	"care-taker/controllers"
	"care-taker/middlewares"

	"github.com/gin-gonic/gin"
)

func EventRoutes(r *gin.RouterGroup) {
	r.GET("/", controllers.GetAllEvents)
	r.POST("/", controllers.CreateEvent)
	r.GET("/:id", controllers.GetEventById)
	r.PUT("/:id", middlewares.IsAdmin, controllers.UpdateEvent)
	r.DELETE("/:id", middlewares.IsAdmin, controllers.DeleteEvent)
	r.POST("/:id/register", middlewares.IsAuthenticated, controllers.RegisterToEvent)
	r.POST("/:id/unregister", middlewares.IsAuthenticated, controllers.RegisterToEvent)
}
