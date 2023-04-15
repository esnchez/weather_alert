package service

import "fmt"

type Notification interface {
	Send() error
}

type NotificationService struct {
}

func NewNotificationService() *NotificationService {
	return &NotificationService{}
}

func (ns *NotificationService) Send() error {
	fmt.Println("An alert message has been sent!!!!")
	return nil
}
