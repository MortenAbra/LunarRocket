package handler

import (
	"context"
	"log"
	"lunar/types"
)

type RocketHandler interface {
	RocketLaunched(ctx context.Context, r types.RocketData) error
	RocketSpeedIncrease(ctx context.Context, r types.RocketData) error
	RocketSpeedDecrease(ctx context.Context, r types.RocketData) error
	RocketExploded(ctx context.Context, r types.RocketData) error
	RocketMissionChanged(ctx context.Context, r types.RocketData) error
}

type RocketHandlerImpl struct {
}

func GetRocketHandlerInstance() RocketHandler {
	return &RocketHandlerImpl{}
}

func (handler *RocketHandlerImpl) RocketLaunched(ctx context.Context, r types.RocketData) error {
	log.Println(r)

	return nil
}

func (handler *RocketHandlerImpl) RocketSpeedIncrease(ctx context.Context, r types.RocketData) error {
	log.Println(r)

	return nil
}

func (handler *RocketHandlerImpl) RocketSpeedDecrease(ctx context.Context, r types.RocketData) error {
	log.Println(r)

	return nil
}

func (handler *RocketHandlerImpl) RocketExploded(ctx context.Context, r types.RocketData) error {
	log.Println(r)

	return nil
}

func (handler *RocketHandlerImpl) RocketMissionChanged(ctx context.Context, r types.RocketData) error {
	log.Println(r)

	return nil
}
