package service

import (
	"errors"

	"github.com/esnchez/weather_alert/domain/weather"
	"github.com/esnchez/weather_alert/domain/weather/memory"
)

var(
	ErrNoRepositoryConfigured = errors.New("Error calling a function with a non injected repo in the service")
	ErrNoNotificationConfigured = errors.New("Error calling a function with a non injected notification system in the service")

)

type WeatherServiceConfig func(as *WeatherService) error

type WeatherService struct {
	repo         weather.Repository
	notification Notification
}

func NewWeatherService(cfgs ...WeatherServiceConfig) (*WeatherService, error) {

	as := &WeatherService{}

	for _, cfg := range cfgs {
		if err := cfg(as); err != nil {
			return nil, err
		}
	}
	return as, nil
}

func WithMemoryWeatherRepository() WeatherServiceConfig {
	wr := memory.NewMemoryRepository()
	return func(as *WeatherService) error {
		as.repo = wr
		return nil
	}
}

func WithNotificationService() WeatherServiceConfig {
	ns := NewNotificationService()
	return func(as *WeatherService) error {
		as.notification = ns
		return nil
	}
}

func (ws *WeatherService) CreateRegistry(wj weather.WeatherJSONResp) error {

	if ws.repo == nil {
		return ErrNoRepositoryConfigured
	}
	if ws.notification == nil {
		return ErrNoRepositoryConfigured
	}

	wr, err := weather.NewWeatherRegistry(wj)
	if err != nil {
		return err
	}


	if wr.GetStateCode() == weather.BadWeather {
		if err := ws.notification.Send(); err != nil {
			return err
		}
	}

	if err := ws.repo.Save(wr); err != nil {
		return err
	}

	return nil
}

func (ws *WeatherService) GetAllRegistries() ([]*weather.WeatherRegistry, error) {
	return ws.repo.GetAll()
}
