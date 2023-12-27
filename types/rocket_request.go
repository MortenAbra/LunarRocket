package types

import (
	"time"
)

type RocketData struct {
	Metadata struct {
		Channel       string    `json:"channel"`
		MessageNumber int       `json:"messageNumber"`
		MessageType   string    `json:"messageType"`
		MessageTime   time.Time `json:"messageTime"`
	} `json:"metadata"`
	MessageContent MessageContent `json:"message"`
}

type MessageContent interface{}

type RocketLaunched struct {
	MessageContent
	Type        string `json:"type"`
	LaunchSpeed int    `json:"launchSpeed"`
	Mission     string `json:"mission"`
}

type RocketSpeedIncrease struct {
	By int `json:"by"`
}

type RocketSpeedDecrease struct {
	By int `json:"by"`
}

type RocketExploded struct {
	Reason int `json:"reason"`
}

type RocketMissionChanged struct {
	NewMission string `json:"newMission"`
}
