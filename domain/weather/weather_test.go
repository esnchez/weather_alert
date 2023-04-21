package weather_test

import (
	"errors"
	"testing"

	"github.com/esnchez/weather_alert/domain/weather"
	"github.com/stretchr/testify/assert"
)

func Test_CreateAllWeatherRegistryTypes(t *testing.T) {

	type testElement struct {
		test                 string
		weatherInfo          *weather.WeatherJSONResp
		expectedErr          error
		expectedWeatherState weather.State
	}

	testElms := []testElement{

		{
			test: "Creating a bad weather registry",
			weatherInfo: &weather.WeatherJSONResp{
				Name: "Barcelona",
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
			expectedErr: nil,
			expectedWeatherState: weather.BadWeather,

		},
		{
			test: "Creating a neutral weather registry",
			weatherInfo: &weather.WeatherJSONResp{
				Name: "Barcelona",
				Main: weather.Main{
					Temperature: 12.3,
					Humidity:    35,
				},

				Wind: weather.Wind{
					Speed: 10.5,
				},
				Clouds: weather.Clouds{
					All: 80,
				},
			},
			expectedErr: nil,
			expectedWeatherState: weather.NeutralWeather,

		},
		{
			test: "Creating a good weather registry",
			weatherInfo: &weather.WeatherJSONResp{
				Name: "Barcelona",
				Main: weather.Main{
					Temperature: 25,
					Humidity:    50,
				},

				Wind: weather.Wind{
					Speed: 5,
				},
				Clouds: weather.Clouds{
					All: 50,
				},
			},
			expectedErr:          nil,
			expectedWeatherState: weather.GoodWeather,
		},
	}

	for _, te := range testElms {
		t.Run(te.test, func(t *testing.T) {
			w, err := weather.NewWeatherRegistry(te.weatherInfo)
			if err != nil {
				t.Fatal(err)
			}

			if !errors.Is(err, te.expectedErr) {
				t.Errorf("expected error %v, got %v", te.expectedErr, err)
			}

			assert.Equal(t, w.StateCode, te.expectedWeatherState)
		})
	}

}

func Test_CreateWeatherRegistryInvalidParams(t *testing.T) {

	type testElement struct {
		test                 string
		weatherInfo          *weather.WeatherJSONResp
		expectedErr          error
		expectedWeatherState weather.State
	}

	testElms := []testElement{

		{
			test: "Creating a bad weather registry, no city name",
			weatherInfo: &weather.WeatherJSONResp{
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
			expectedErr: weather.ErrInvalidParams,

		},
		{
			test: "Creating a weather registry without values",
			weatherInfo: &weather.WeatherJSONResp{
				Name: "Barcelona",
			},
			expectedErr: weather.ErrInvalidParams,

		},
	}

	for _, te := range testElms {
		t.Run(te.test, func(t *testing.T) {
			_, err := weather.NewWeatherRegistry(te.weatherInfo)

			if !errors.Is(err, te.expectedErr) {
				t.Errorf("expected error %v, got %v", te.expectedErr, err)
			}

		})
	}

}

