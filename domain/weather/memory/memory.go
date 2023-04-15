package memory

import (
	"github.com/esnchez/weather_alert/domain/weather"
)

// var wrList = []*weather.WeatherRegistry{}

type MemoryRepository struct {
	list []*weather.WeatherRegistry
}

func NewMemoryRepository() *MemoryRepository {

	return &MemoryRepository{
		list: make([]*weather.WeatherRegistry, 0),
	}

}

func (mr *MemoryRepository) GetAll() ([]*weather.WeatherRegistry, error) {

	return mr.list, nil
}

func (mr *MemoryRepository) Save(wr *weather.WeatherRegistry) error {
	mr.list = append(mr.list, wr)

	return nil
}
