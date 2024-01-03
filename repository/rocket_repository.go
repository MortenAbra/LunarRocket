package repository

import (
	"context"
	"lunar/connectors"
	"lunar/types"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type RocketRepository interface {
	CreateRocket(ctx context.Context, data types.RocketData) error
	UpdateSpeedIncrease(ctx context.Context, data types.RocketData) error
	UpdateSpeedDecrease(ctx context.Context, data types.RocketData) error
	UpdateRocketStatus(ctx context.Context, data types.RocketData) error
	UpdateRocketMission(ctx context.Context, data types.RocketData) error
}

type RocketRepositoryImpl struct {
	db *gorm.DB
}

func GetRocketRepositoryInstance() RocketRepository {
	return &RocketRepositoryImpl{
		db: connectors.GetDB(),
	}
}

func (repo RocketRepositoryImpl) CreateRocket(ctx context.Context, data types.RocketData) error {
	rocket := types.RocketModel{
		Channel:       data.Metadata.Channel,
		MessageNumber: data.Metadata.MessageNumber,
		MessageType:   data.Metadata.MessageType,
		Type:          data.MessageContent.Type,
		LaunchSpeed:   data.MessageContent.LaunchSpeed,
		Mission:       data.MessageContent.Mission,
		Speed:         data.MessageContent.LaunchSpeed,
		Status:        data.MessageContent.Reason,
		NewMission:    data.MessageContent.NewMission,
	}
	result := repo.db.Clauses(clause.OnConflict{DoNothing: true}).Create(&rocket)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repo RocketRepositoryImpl) UpdateSpeedIncrease(ctx context.Context, data types.RocketData) error {
	result := repo.db.Table("rockets").
		Where("channel = ?", data.Metadata.Channel).
		Updates(map[string]interface{}{
			"speed":          gorm.Expr("speed + ?", data.MessageContent.By),
			"message_number": data.Metadata.MessageNumber,
			"message_time":   data.Metadata.MessageTime,
			"message_type":   data.Metadata.MessageType,
		})

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo RocketRepositoryImpl) UpdateSpeedDecrease(ctx context.Context, data types.RocketData) error {
	result := repo.db.Table("rockets").
		Where("channel = ?", data.Metadata.Channel).
		Updates(map[string]interface{}{
			"speed":          gorm.Expr("speed - ?", data.MessageContent.By),
			"message_number": data.Metadata.MessageNumber,
			"message_time":   data.Metadata.MessageTime,
			"message_type":   data.Metadata.MessageType,
		})

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo RocketRepositoryImpl) UpdateRocketStatus(ctx context.Context, data types.RocketData) error {

	result := repo.db.Table("rockets").Where("channel = ?", data.Metadata.Channel).Updates(&types.RocketModel{
		MessageNumber: data.Metadata.MessageNumber,
		Status:        data.MessageContent.Reason,
		MessageTime:   data.Metadata.MessageTime,
		MessageType:   data.Metadata.MessageType,
	})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo RocketRepositoryImpl) UpdateRocketMission(ctx context.Context, data types.RocketData) error {

	result := repo.db.Table("rockets").Where("channel = ?", data.Metadata.Channel).Updates(&types.RocketModel{
		MessageNumber: data.Metadata.MessageNumber,
		NewMission:    data.MessageContent.NewMission,
		MessageTime:   data.Metadata.MessageTime,
		MessageType:   data.Metadata.MessageType,
	})
	if result.Error != nil {
		return result.Error
	}
	return nil
}