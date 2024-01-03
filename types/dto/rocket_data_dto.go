package dto

import (
	"time"
)

type RocketLaunchedDto struct {
	Metadata struct {
		Channel       string    `json:"channel"`
		MessageNumber int       `json:"messageNumber"`
		MessageType   string    `json:"messageType"`
		MessageTime   time.Time `json:"messageTime"`
	} `json:"metadata"`
	MessageContent MessageContent `json:"message"`
}

type MessageContent struct {
	Type        string `json:"type,omitempty"`
	LaunchSpeed int    `json:"launchSpeed,omitempty"`
	Mission     string `json:"mission,omitempty"`
	By          int    `json:"by,omitempty"`
	Reason      string `json:"reason,omitempty"`
	NewMission  string `json:"newMission,omitempty"`
}
