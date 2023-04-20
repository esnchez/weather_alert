package weather

import (
	"errors"
)


var (
	ErrFailedToCreateRegistry = errors.New("Failed to create registry and add it to the repository")
)

type Repository interface {
	GetAll() ([]*WeatherRegistry, error)
	Save(*WeatherRegistry) error
}
