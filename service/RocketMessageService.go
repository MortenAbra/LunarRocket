package service

import (
	"lunar/types"
	"sync"
)

type RocketMessageService interface {
	ProcessMessage(data types.RocketData) ([]types.RocketData, error)
}

type RocketMessageServiceImpl struct {
	messageQueues   map[string]map[int]types.RocketData
	expectedNumbers map[string]int
	mu              sync.Mutex
}

func NewRocketMessageService() RocketMessageService {
	return &RocketMessageServiceImpl{
		messageQueues:   make(map[string]map[int]types.RocketData),
		expectedNumbers: make(map[string]int),
	}
}

func (s *RocketMessageServiceImpl) ProcessMessage(data types.RocketData) ([]types.RocketData, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	channel := data.Metadata.Channel
	if _, exists := s.messageQueues[channel]; !exists {
		s.messageQueues[channel] = make(map[int]types.RocketData)
		s.expectedNumbers[channel] = 1
	}

	var processedMessages []types.RocketData

	if data.Metadata.MessageNumber == s.expectedNumbers[channel] {
		processedMessages = append(processedMessages, data)
		s.expectedNumbers[channel]++

		for {
			if nextData, ok := s.messageQueues[channel][s.expectedNumbers[channel]]; ok {
				processedMessages = append(processedMessages, nextData)
				delete(s.messageQueues[channel], s.expectedNumbers[channel])
				s.expectedNumbers[channel]++
			} else {
				break
			}
		}
	} else {
		s.messageQueues[channel][data.Metadata.MessageNumber] = data
	}

	return processedMessages, nil
}
