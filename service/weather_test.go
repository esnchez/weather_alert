package service_test

import (
	"fmt"
	"testing"

	"github.com/esnchez/weather_alert/domain/weather"
	"github.com/esnchez/weather_alert/service"
)

func Test_CreateRegistry(t *testing.T) {

	ws, err := service.NewWeatherService(
		service.WithMemoryWeatherRepository(),
		service.WithNotificationService(),
	)
	if err != nil {
		t.Fatal(err)
	}

	jsonMock := weather.WeatherJSONResp{
		Name: "Barcelona",
		Weather: []weather.Weather{
			{
				Description: "cloudy day",
			},
		},
		Main: weather.Main{
			Temperature: 23.54,
			Humidity:    35,
		},

		Wind: weather.Wind{
			Speed: 50.34,
		},
		Clouds: weather.Clouds{
			All: 80,
		},
	}

	if err := ws.CreateRegistry(jsonMock); err != nil {
		t.Fatal(err)
	}

	fromdb, err := ws.GetAllRegistries()

	for _, v := range fromdb {
		fmt.Println("value retrieved:", v)
	}

	fmt.Println("Test ended! ")

}
