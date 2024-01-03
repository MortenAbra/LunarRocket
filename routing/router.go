package routing

import "github.com/gin-gonic/gin"

// Defines the router instance with the different groups of requests
func RouterInstance(ctx *gin.Context) *gin.Engine {
	r := gin.Default()

	v1 := r.Group("v1/")
	addRocketRoutes(v1)

	return r
}
