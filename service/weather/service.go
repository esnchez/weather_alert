package service

import (
	"log"

	"github.com/esnchez/weather_alert/domain/weather"
)

func (ws *WeatherService) CreateRegistry(wj *weather.WeatherJSONResp) error {

	if ws.repo == nil || ws.alert == nil{
		return ErrNoConfiguredService
	}

	wr, err := weather.NewWeatherRegistry(wj)
	if err != nil {
		log.Printf("Error creating the registry")
		return err
	}

	//Send alert point
	if wr.StateCode == weather.BadWeather {
		if err := ws.alert.Send(wr.CityName); err != nil {
			log.Printf("Error sending the alert")
			return err
		}
	}

	if err := ws.repo.Save(wr); err != nil {
		log.Printf("Error saving registry in the repo")
		return err
	}

	return nil
}

func (ws *WeatherService) GetAllRegistries() ([]*weather.WeatherRegistry, error) {

	if ws.repo == nil || ws.alert == nil{
		return nil, ErrNoConfiguredService
	}

	return ws.repo.GetAll()
}