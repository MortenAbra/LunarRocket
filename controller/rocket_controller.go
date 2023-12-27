package controller

import (
	"lunar/handler"
	"lunar/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RocketController interface {
	GetMessages(ctx *gin.Context)
}

type RocketControllerImpl struct {
	RocketHandler handler.RocketHandler
}

func GetRocketControllerInstance() RocketController {
	return &RocketControllerImpl{
		RocketHandler: handler.GetRocketHandlerInstance(),
	}
}

func (controller *RocketControllerImpl) GetMessages(ctx *gin.Context) {
	content, exists := ctx.Get("messageContent")
	if !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Message content not found"})
		return
	}

	msg, ok := content.(types.RocketData)
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid message content type"})
		return
	}

	switch msg.MessageContent.(type) {
	case types.RocketLaunched:
		err := controller.RocketHandler.RocketLaunched(ctx, msg)
		if err != nil {
			return
		}
		ctx.JSON(200, nil)
	case types.RocketSpeedIncrease:
		err := controller.RocketHandler.RocketSpeedIncrease(ctx, msg)
		if err != nil {
			return
		}
		ctx.JSON(200, nil)
	case types.RocketSpeedDecrease:
		err := controller.RocketHandler.RocketSpeedDecrease(ctx, msg)
		if err != nil {
			return
		}
		ctx.JSON(200, nil)
	case types.RocketExploded:
		err := controller.RocketHandler.RocketExploded(ctx, msg)
		if err != nil {
			return
		}
		ctx.JSON(200, nil)
	case types.RocketMissionChanged:
		err := controller.RocketHandler.RocketMissionChanged(ctx, msg)
		if err != nil {
			return
		}
		ctx.JSON(200, nil)
	// Add cases for other message types...
	default:
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid message content type"})
		return
	}
}
