package weather

import (
	"errors"
	"log"

	"github.com/google/uuid"
)

var (
	ErrInvalidParams = errors.New("Weather registry creation failed because of incomplete information")
)

// Principal entity in our business domain
type WeatherRegistry struct {
	Id          string  `json:"id"`
	CityName    string  `json:"cityname"`
	Temperature float64 `json:"temperature"`
	StateCode   State   `json:"statecode"`
	StateDesc   string  `json:"description"`
}

// Enum-like State has its pre-defined values
type State int

const (
	UndefinedWeather State = iota
	BadWeather
	NeutralWeather
	GoodWeather
)

// Limits in 4 weather parameters set to calculate Weather State. They are stored in memory but could be taken from a configuration file/db
//Temperature: Represented in ºC
//Humidity: Percentage
//Wind Speed: ms/s
//Cloud: Percentage

var (
	tempLimit            = 15
	humLimit             = 80
	windSpeedLimit       = 12
	cloudPercentageLimit = 70

	limitArray      = []int{tempLimit, humLimit, windSpeedLimit, cloudPercentageLimit}
	isValueMaxArray = []bool{false, true, true, true}
)

// Factory to create new weather registries. Applies certain rules to define the final weather state.
func NewWeatherRegistry(w *WeatherJSONResp) (*WeatherRegistry, error) {

	if w.Name == "" {
		return nil, ErrInvalidParams
	}

	weatherState, desc := calculateWeatherState(int(w.Main.Temperature), w.Main.Humidity, int(w.Wind.Speed), w.Clouds.All)
	log.Println("Created weather state and description: ", weatherState, desc)
	if weatherState == UndefinedWeather {
		return nil, ErrInvalidParams
	}

	return &WeatherRegistry{
		Id:          uuid.New().String(),
		CityName:    w.Name,
		Temperature: w.Main.Temperature,
		StateCode:   weatherState,
		StateDesc:   desc,
	}, nil
}

// Helper function to retrieve an explanatory string depending on State value
func getStateDesc(s State) string {
	if s == 0 {
		return "Undefined weather"
	}
	if s == 1 {
		return "Bad weather"
	}
	if s == 3 {
		return "Good weather"
	}
	return "Neutral weather"
}

// Returns the final weather State (int) based on the sumatory of each weather category ponderation
func calculateWeatherState(temp, hum, ws, cp int) (State, string) {
	var weatherValue int

	weaValArray := [4]int{temp, hum, ws, cp}

	for i := range weaValArray {
		val := calculateVal(isValueMaxArray[i], weaValArray[i], limitArray[i])
		weatherValue += val
	}

	if weatherValue == 0 {
		return UndefinedWeather, getStateDesc(UndefinedWeather)
	}
	if weatherValue >= 7 {
		return BadWeather, getStateDesc(BadWeather)
	}
	if weatherValue <= 4 {
		return GoodWeather, getStateDesc(GoodWeather)
	}
	return NeutralWeather, getStateDesc(NeutralWeather)
}

// This function returns a numeric value depending on the operation and the comparison between the limit established and the weather category value
func calculateVal(isValueMax bool, value, limit int) int {

	switch isValueMax {
	// humidity, wind speed and cloud percentage evaluation
	case true:
		if value == 0 {
			return 0
		}
		if value > limit {
			return 2
		}
		return 1
	// temperature evaluation
	case false:
		if value == 0 {
			return 0
		}
		if value < limit-10 {
			return 4
		}
		if value < limit-5 {
			return 3
		}
		if value < limit {
			return 2
		}
		return 1
	default:
		panic("Error: unexpected boolean value")
	}

}
