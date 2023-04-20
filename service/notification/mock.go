package notification

import (
	"log"
)

type MockService struct {
}

func NewMockService() *MockService{
	return &MockService{}
}

func (ms *MockService) Send(cityName string) error {
	log.Printf("ALERT SENT! Bad weather in %v!", cityName)
	return nil
}
