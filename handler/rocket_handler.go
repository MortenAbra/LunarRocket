package handler

import (
	"context"
	"fmt"
	"lunar/repository"
	"lunar/service"
	"lunar/types"
)

type RocketHandler interface {
	ProcessAndHandle(ctx context.Context, data types.RocketData) error
	GetRockets(ctx context.Context) ([]types.RocketModel, error)
}

type RocketHandlerImpl struct {
	RocketRepository     repository.RocketRepository
	RocketMessageService service.RocketMessageService
}

func GetRocketHandlerInstance() RocketHandler {
	return &RocketHandlerImpl{
		RocketRepository:     repository.GetRocketRepositoryInstance(),
		RocketMessageService: service.NewRocketMessageService(),
	}
}

// Function for handling processed data by RocketMessageService and distributing them to correct handler functions
func (handler *RocketHandlerImpl) ProcessAndHandle(ctx context.Context, data types.RocketData) error {
	processedMessages, err := handler.RocketMessageService.ProcessMessage(data)
	if err != nil {
		return err
	}

	for _, msg := range processedMessages {
		switch msg.Metadata.MessageType {
		case "RocketLaunched":
			err = handler.launchRocket(ctx, msg)
		case "RocketSpeedIncreased":
			err = handler.increaseSpeed(ctx, msg)
		case "RocketSpeedDecreased":
			err = handler.decreaseSpeed(ctx, msg)
		case "RocketExploded":
			err = handler.explodeRocket(ctx, msg)
		case "RocketMissionChanged":
			err = handler.changeMission(ctx, msg)
		default:
			err = fmt.Errorf("unknown message type: %s", msg.Metadata.MessageType)
		}

		if err != nil {
			return err
		}
	}

	return nil
}

// Functions for submitting data to repository for updating rockets for specific actions
func (handler *RocketHandlerImpl) launchRocket(ctx context.Context, r types.RocketData) error {
	return handler.RocketRepository.CreateRocket(ctx, r)
}

func (handler *RocketHandlerImpl) increaseSpeed(ctx context.Context, r types.RocketData) error {
	return handler.RocketRepository.UpdateSpeedIncrease(ctx, r)
}

func (handler *RocketHandlerImpl) decreaseSpeed(ctx context.Context, r types.RocketData) error {
	return handler.RocketRepository.UpdateSpeedDecrease(ctx, r)
}

func (handler *RocketHandlerImpl) explodeRocket(ctx context.Context, r types.RocketData) error {
	return handler.RocketRepository.UpdateRocketStatus(ctx, r)
}

func (handler *RocketHandlerImpl) changeMission(ctx context.Context, r types.RocketData) error {
	return handler.RocketRepository.UpdateRocketMission(ctx, r)
}

// Function for getting rockets from repository
func (handler *RocketHandlerImpl) GetRockets(ctx context.Context) ([]types.RocketModel, error) {
	rockets, err := handler.RocketRepository.GetRockets(ctx)
	if err != nil {
		return rockets, err
	}

	return rockets, nil
}
