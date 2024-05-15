package routes

import (
	"care-taker/controllers"

	"github.com/gin-gonic/gin"
)

func EventRoutes(r *gin.RouterGroup) {
	r.GET("/", controllers.GetAllEvents)
	r.POST("/", controllers.CreateEvent)
	r.GET("/:id", controllers.GetEventById)
	r.PUT("/:id", controllers.UpdateEvent)
	r.DELETE("/:id", controllers.DeleteEvent)
}
