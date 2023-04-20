package notification

import (
	"errors"
)

var (
	ErrSendingAlert = errors.New("Error: Alert service has failed sending the alert")
)

type Notification interface {
	Send(cityName string) error
}
