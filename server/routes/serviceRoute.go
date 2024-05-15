package routes

import (
	"care-taker/controllers"

	"github.com/gin-gonic/gin"
)

func ServiceRoutes(r *gin.RouterGroup) {
	r.GET("/", controllers.GetAllServices)
	r.POST("/", controllers.CreateService)
	r.GET("/:id", controllers.GetServiceById)
	r.PUT("/:id", controllers.UpdateService)
	r.DELETE("/:id", controllers.DeleteService)
}
