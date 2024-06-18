package routes

import (
	"care-taker/controllers"
	"care-taker/middlewares"

	"github.com/gin-gonic/gin"
)

func ServiceRoutes(r *gin.RouterGroup) {
	r.GET("/", controllers.GetAllServices)
	r.POST("/", controllers.CreateService)
	r.GET("/:id", controllers.GetServiceById)
	r.PUT("/:id", middlewares.IsAdmin, controllers.UpdateService)
	r.DELETE("/:id", middlewares.IsAdmin, controllers.DeleteService)
}
