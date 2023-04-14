package weather

import (
	"github.com/google/uuid"
)

type WeatherInfo struct {
	CityName        string
	Description     string
	Temperature     float64
	Humidity        int
	WindSpeed       int
	CloudPercentage int
}

type WeatherRegistry struct {
	id        uuid.UUID
	values    WeatherInfo
	stateCode State
	stateName string
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
	tempLimit      = 15
	humLimit       = 80
	windSpeedLimit = 8
	cloudLimit     = 70

	limitArray          = []int{tempLimit, humLimit, windSpeedLimit, cloudLimit}
	operationIsMaxArray = []bool{false, true, true, true}
)

// Factory to create new weather registries that apply certain rules to define the final weather state.
func NewWeatherRegistry(wr WeatherInfo) *WeatherRegistry {

	ws := calculateWeatherState(wr.Temperature, wr.Humidity, wr.WindSpeed, wr.CloudPercentage)
	name := getWeatherStateName(ws)

	return &WeatherRegistry{
		id:        uuid.New(),
		values:    wr,
		stateCode: ws,
		stateName: name,
	}

}

func getWeatherStateName(s State) string {
	if s == 1 {
		return "Bad weather"
	}
	if s == 3 {
		return "Good weather"
	}
	return "Neutral weather"
}

func calculateWeatherState(temp float64, hum, ws, cp int) State {
	var weatherValue int

	valuesArray := [4]int{int(temp), hum, ws, cp}

	for i, _ := range valuesArray {
		val := calculateVal(operationIsMaxArray[i], valuesArray[i], limitArray[i])
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

// This function sets
func calculateVal(isMax bool, value, limit int) int {
	switch isMax {
	case true:
		if value > limit {
			return 2
		}
		return 0
	case false:
		if value < limit {
			return 2
		}
		return 0
	default:
		panic("Error with unexpected value")
	}

}
