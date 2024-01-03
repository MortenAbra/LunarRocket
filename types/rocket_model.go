package types

import "time"

type RocketModel struct {
	Channel       string `gorm:"primaryKey"`
	MessageNumber int    `gorm:"index"`
	MessageType   string
	MessageTime   time.Time
	Type          string
	LaunchSpeed   int
	Mission       string
	Speed         int
	Status        string
	NewMission    string
}

func (RocketModel) TableName() string {
	return "rockets"
}
