package service

import (
	"database/sql"
	"errors"

	"github.com/esnchez/weather_alert/domain/weather"
	"github.com/esnchez/weather_alert/domain/weather/memory"
	"github.com/esnchez/weather_alert/domain/weather/mysql"

	"github.com/esnchez/weather_alert/service/notification"
)

var (
	ErrNoConfiguredService = errors.New("Error: Service has been initialized without injecting dependencies")
)

type WeatherServiceConfig func(as *WeatherService) error

type WeatherService struct {
	repo  weather.Repository
	alert notification.Notification
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

func WithMySQLWeatherRepository(db *sql.DB) WeatherServiceConfig {
	wr := mysql.NewMySQLRepository(db)
	return func(as *WeatherService) error {
		as.repo = wr
		return nil
	}
}

func WithTwilioService(cfg *notification.TwilioConfig) WeatherServiceConfig {
	ns := notification.NewTwilioService(cfg)
	return func(as *WeatherService) error {
		as.alert = ns
		return nil
	}
}

func WithMockNotificationService() WeatherServiceConfig {
	ns := notification.NewMockService()
	return func(as *WeatherService) error {
		as.alert = ns
		return nil
	}
}

