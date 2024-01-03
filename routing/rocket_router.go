package routing

import (
	"lunar/controller"
	"lunar/middleware"

	"github.com/gin-gonic/gin"
)

// Creates rocket routes
func addRocketRoutes(r *gin.RouterGroup) {
	rocketController := controller.GetRocketControllerInstance()

	r.Use(gin.Recovery())


	// Rocket group containg POST and GET requests
	rocketGroup := r.Group("/rockets")
	{
		rocketGroup.POST("/messages", middleware.MessageDispatchMiddleware(), rocketController.GetMessages)
		rocketGroup.GET("", rocketController.GetRockets)
	}
}
