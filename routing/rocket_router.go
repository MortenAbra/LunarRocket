package routing

import (
	"lunar/controller"
	"lunar/middleware"

	"github.com/gin-gonic/gin"
)

func addRocketRoutes(r *gin.RouterGroup) {
	rocketController := controller.GetRocketControllerInstance()

	r.Use(gin.Recovery())

	rocketGroup := r.Group("/rockets")
	{
		rocketGroup.POST("/messages", middleware.MessageDispatchMiddleware(), rocketController.GetMessages)
		rocketGroup.GET("", rocketController.GetRockets)
	}
}
