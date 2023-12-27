package middleware

import (
	"encoding/json"
	"lunar/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

func MessageDispatchMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var msg types.RocketData
		if err := c.BindJSON(&msg); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			c.Abort()
			return
		}


		// Add logic here to handle different message types
		switch msg.Metadata.MessageType {
		case "RocketLaunched":
			var m types.RocketLaunched
			if err := json.Unmarshal(messageJSON, &m); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				c.Abort()
				return
			}
			msg.MessageContent = m
			c.Set("messageContent", msg)
		case "RocketSpeedIncrease":
			var m types.RocketSpeedIncrease
			if err := json.Unmarshal(messageJSON, &m); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				c.Abort()
				return
			}
			msg.MessageContent = m
			c.Set("messageContent", msg)
		case "RocketSpeedDecrease":
			var m types.RocketSpeedDecrease
			if err := json.Unmarshal(msg.MessageContent.(json.RawMessage), &m); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				c.Abort()
				return
			}
			msg.MessageContent = m
			c.Set("messageContent", msg)
		case "RocketExploded":
			var m types.RocketExploded
			if err := json.Unmarshal(msg.MessageContent.(json.RawMessage), &m); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				c.Abort()
				return
			}
			msg.MessageContent = m
			c.Set("messageContent", msg)
		case "RocketMissionChanged":
			var m types.RocketMissionChanged
			if err := json.Unmarshal(msg.MessageContent.(json.RawMessage), &m); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				c.Abort()
				return
			}
			msg.MessageContent = m
			c.Set("messageContent", msg)
		default:
			c.JSON(http.StatusBadRequest, gin.H{"error": "Unknown message type"})
			c.Abort()
			return
		}

		c.Next()
	}
}
