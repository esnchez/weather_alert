package weather

import (
	"errors"

	"github.com/google/uuid"
)

type WeatherRegistry struct {
	id              uuid.UUID
	cityName        string
	description     string
	temperature     float64
	humidity        int
	windSpeed       float64
	cloudPercentage int
	stateCode       State
	stateName       string
}

// Enum-like State, that has its pre-defined values
type State int

const (
	UndefinedWeather State = iota
	BadWeather
	NeutralWeather
	GoodWeather
)

// Limits set to calculate Weather State. They are stored in memory but could be taken from a configuration file
var (
	tempLimit            = 15
	humLimit             = 80
	windSpeedLimit       = 12
	cloudPercentageLimit = 70

	limitArray      = []int{tempLimit, humLimit, windSpeedLimit, cloudPercentageLimit}
	isValueMaxArray = []bool{false, true, true, true}

	ErrMissingValues = errors.New("Weather registry has to include weather valid values")
)

// Factory to create new weather registries that apply certain rules to define the final weather state.
func NewWeatherRegistry(w WeatherJSONResp) (*WeatherRegistry, error) {

	//TEST HOW TO SEND AN INVALID WEATHER INFO

	weatherState := calculateWeatherState(int(w.Main.Temperature), w.Main.Humidity, int(w.Wind.Speed), w.Clouds.All)
	name := getStateName(weatherState)

	return &WeatherRegistry{
		id:              uuid.New(),
		cityName:        w.Name,
		description:     w.Weather[0].Description,
		temperature:     w.Main.Temperature,
		humidity:        w.Main.Humidity,
		windSpeed:       w.Wind.Speed,
		cloudPercentage: w.Clouds.All,
		stateCode:       weatherState,
		stateName:       name,
	}, nil
}

func (wr *WeatherRegistry) GetState() string {
	return wr.stateName
}

func (wr *WeatherRegistry) GetStateCode() State {
	return wr.stateCode
}

// Helper function to retrieve an explanatory string depending on State value
func getStateName(s State) string {
	if s == 1 {
		return "Bad weather"
	}
	if s == 3 {
		return "Good weather"
	}
	return "Neutral weather"
}

// Returns the final weather State (int) based on the sumatory of each weather category ponderation
func calculateWeatherState(temp, hum, ws, cp int) State {
	var weatherValue int

	weaValArray := [4]int{temp, hum, ws, cp}

	for i := range weaValArray {
		val := calculateVal(isValueMaxArray[i], weaValArray[i], limitArray[i])
		weatherValue += val
	}

	if weatherValue >= 6 {
		return BadWeather
	}
	if weatherValue <= 2 {
		return GoodWeather
	}
	return NeutralWeather
}

// This function returns a numeric value depending on the operation and the comparison between the limit established and the weather category value
func calculateVal(isValueMax bool, value, limit int) int {
	switch isValueMax {
	case true:
		if value > limit {
			return 2
		}
		return 0
	// temperature evaluation
	case false:
		if value < limit-10 {
			return 3
		}
		if value < limit-5 {
			return 2
		}
		if value < limit {
			return 1
		}
		return 0
	default:
		panic("Error: unexpected boolean value")
	}

}
