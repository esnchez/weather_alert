package service_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/esnchez/weather_alert/domain/weather"
	"github.com/esnchez/weather_alert/service/notification"
	"github.com/esnchez/weather_alert/service/weather"
	"github.com/stretchr/testify/assert"
)

func Test_CreateWeatherRegistryWithIncompleteServiceParams(t *testing.T) {

	weatherInfo := &weather.WeatherJSONResp{
			Name: "Barcelona",
			Weather: []weather.Weather{
				{
					Description: "cloudy and windy day",
				},
			},
			Main: weather.Main{
				Temperature: 4.5,
				Humidity:    90,
			},

			Wind: weather.Wind{
				Speed: 20,
			},
			Clouds: weather.Clouds{
				All: 50,
			},
	}

	//Service 1: no notification system
	ws1, err := service.NewWeatherService(
		service.WithMemoryWeatherRepository(),
	) 	
	if err != nil {
		t.Fatal(err)
	}
	
	if err := ws1.CreateRegistry(weatherInfo); err != nil {
		
		if !errors.Is(err, service.ErrNoConfiguredService) {
			t.Errorf("expected error %v, got %v", service.ErrNoConfiguredService, err)
		}
	}
	
	//Service 1: no repository
	ws2, err := service.NewWeatherService(
		service.WithMockNotificationService(),
	)
	if err != nil {
		t.Fatal(err)
	}

	if err := ws2.CreateRegistry(weatherInfo); err != nil {
		
		if !errors.Is(err, service.ErrNoConfiguredService) {
			t.Errorf("expected error %v, got %v", service.ErrNoConfiguredService, err)
		}
	}


	fmt.Println("Test ended! ")
}


func Test_CreateBadWeatherRegistryIntoMemoryRepositoryAndAlertMockService(t *testing.T) {

	type testElement struct {
		test                 string
		weatherInfo          *weather.WeatherJSONResp
		expectedErr          error
		expectedWeatherState weather.State
	}

	testElms := []testElement{

		{
			test: "Creating a bad weather registry into memory repository",
			weatherInfo: &weather.WeatherJSONResp{
				Name: "Barcelona",
				Weather: []weather.Weather{
					{
						Description: "cloudy and windy day",
					},
				},
				Main: weather.Main{
					Temperature: 4.5,
					Humidity:    90,
				},

				Wind: weather.Wind{
					Speed: 20,
				},
				Clouds: weather.Clouds{
					All: 50,
				},
			},
			expectedErr:          nil,
			expectedWeatherState: weather.BadWeather,
		},
	}

	for _, te := range testElms {
		t.Run(te.test, func(t *testing.T) {

			ws, err := service.NewWeatherService(
				service.WithMemoryWeatherRepository(),
				service.WithMockNotificationService(),

			)
			if err != nil {
				t.Fatal(err)
			}

			if err := ws.CreateRegistry(te.weatherInfo); err != nil {
				t.Fatal(err)
			}
			
			if !errors.Is(err, te.expectedErr) {
				t.Errorf("expected error %v, got %v", te.expectedErr, err)
			}
			
			list, err := ws.GetAllRegistries()
			if err != nil {
				t.Fatal(err)
			}

			for _, v := range list {
				assert.Equal(t, v.StateCode, te.expectedWeatherState)
				assert.Equal(t, v.CityName, te.weatherInfo.Name)
				assert.Equal(t, v.Temperature, te.weatherInfo.Main.Temperature)
				fmt.Printf("City name: %v, Weather Status code: %v, Temperature registered: %v and description: %v \n", v.CityName, v.StateCode, v.Temperature, v.StateDesc)
			}
		})
	}
	fmt.Println("Test ended! ")
}


func Test_CreateBadWeatherRegistryWithTwilio(t *testing.T) {

	hardcodedCfg := &notification.TwilioConfig{
		AccountSid:  "AC282d5e9f31a619f08496ccb9150460d5",
		Auth:        "4904aa959e73f11f2193fb20a90598e4",
		AlertPhone:  "0",
		PersPhone:   "+34650870690",
		TwilioPhone: "+13187318414",
		AlertMsg:    "Bad weather alert in: ",
	}

	ws, err := service.NewWeatherService(
		service.WithMemoryWeatherRepository(),
		service.WithTwilioService(hardcodedCfg),
	)
	if err != nil {
		t.Fatal(err)
	}

	jsonMock := &weather.WeatherJSONResp{
		Name: "Berlin",
		Weather: []weather.Weather{
			{
				Description: "cloudy and windy day",
			},
		},
		Main: weather.Main{
			Temperature: 4.5,
			Humidity:    65,
		},

		Wind: weather.Wind{
			Speed: 30.34,
		},
		Clouds: weather.Clouds{
			All: 80,
		},
	}

	if err := ws.CreateRegistry(jsonMock); err != nil {
		t.Fatal(err)
	}

	mem, err := ws.GetAllRegistries()
	if err != nil {
		t.Fatal(err)
	}

	for _, v := range mem {
		fmt.Printf("City name: %v, Weather Status code: %v, Temperature registered: %v and description: %v \n", v.CityName, v.StateCode, v.Temperature, v.StateDesc)
	}

	fmt.Println("Test ended! ")
}
