package controller

import (
	"lunar/handler"
	"lunar/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RocketController interface {
	GetMessages(ctx *gin.Context)
	GetRockets(ctx *gin.Context)
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

	// Saves the MessageContent as msg
	msg, ok := content.(types.RocketData)
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid message content type"})
		return
	}

	err := controller.RocketHandler.ProcessAndHandle(ctx, msg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred during processing of data"})
	}

	ctx.JSON(http.StatusAccepted, gin.H{})
}

func (controller *RocketControllerImpl) GetRockets(ctx *gin.Context) {
	rockets, err := controller.RocketHandler.GetRockets(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "with retrieving rockets"})
	}

	rocketResponse := types.MapToRocketResponse(rockets)

	ctx.JSON(http.StatusOK, rocketResponse)
}
