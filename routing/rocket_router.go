package routing

import (
	"lunar/controller"
	"lunar/middleware"

	"github.com/gin-gonic/gin"
)

func addRocketRoutes(r *gin.RouterGroup) {
	rocketController := controller.GetRocketControllerInstance()

	r.Use(gin.Recovery(), middleware.MessageDispatchMiddleware())

	rocketGroup := r.Group("/rockets")
	{
		rocketGroup.POST("/messages", rocketController.GetMessages)
	}
}
