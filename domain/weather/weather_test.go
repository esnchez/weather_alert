package weather_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/esnchez/weather_alert/domain/weather"
)

func Test_CreateWeatherRegistry(t *testing.T) {

	type testElement struct {
		test        string
		weatherInfo weather.WeatherJSONResp
		expectedErr error
	}

	testElms := []testElement{
		// {
		// 	test:        "Creating a weather registry with incomplete parameters",
		// 	wInfo:       weather.WeatherInfo{
		// 		CityName:        "Barcelona",
		// 		Description:     "Whatever desc",
		// 		Humidity:        50,
		// 		WindSpeed:       10,
		// 		CloudPercentage: 20,
		// 	},
		// 	expectedErr: weather.ErrMissingValues,
		// },
		{
			test: "Creating a bad weather registry",
			weatherInfo: weather.WeatherJSONResp{
				Name: "Barcelona",
				Weather: []weather.Weather{
					{
						Description: "cloudy day",
					},
				},
				Main: weather.Main{
					Temperature: 20.54,
					Humidity:    35,
				},

				Wind: weather.Wind{
					Speed: 10.34,
				},
				Clouds: weather.Clouds{
					All: 50,
				},
			},
			expectedErr: nil,
		},
		{
			test: "Creating a neutral weather registry",
			weatherInfo: weather.WeatherJSONResp{
				Name: "Barcelona",
				Weather: []weather.Weather{
					{
						Description: "cloudy day",
					},
				},
				Main: weather.Main{
					Temperature: 20.54,
					Humidity:    35,
				},

				Wind: weather.Wind{
					Speed: 10.34,
				},
				Clouds: weather.Clouds{
					All: 50,
				},
			},
			expectedErr: nil,
		},
		{
			test: "Creating a good weather registry",
			weatherInfo: weather.WeatherJSONResp{
				Name: "Barcelona",
				Weather: []weather.Weather{
					{
						Description: "cloudy day",
					},
				},
				Main: weather.Main{
					Temperature: 20.54,
					Humidity:    35,
				},

				Wind: weather.Wind{
					Speed: 10.34,
				},
				Clouds: weather.Clouds{
					All: 50,
				},
			},
			expectedErr: nil,
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

	w := weather.WeatherJSONResp{
		Name: "Barcelona",
		Weather: []weather.Weather{
			{
				Description: "cloudy day",
			},
		},
		Main: weather.Main{
			Temperature: 20.54,
			Humidity:    35,
		},

		Wind: weather.Wind{
			Speed: 10.34,
		},
		Clouds: weather.Clouds{
			All: 50,
		},
	}

	wr, err := weather.NewWeatherRegistry(w)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("State of the weather today:", wr.GetState())
	fmt.Println(wr.GetStateCode())

}
