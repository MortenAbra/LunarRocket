package middleware

import (
	"lunar/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

func MessageDispatchMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Binds the rocket data to a type and sets it in the context
		var msg types.RocketData
		if err := c.BindJSON(&msg); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		c.Set("messageContent", msg)

		c.Next()
	}
}
