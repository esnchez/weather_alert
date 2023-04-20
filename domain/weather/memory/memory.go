package memory

import (
	"fmt"

	"github.com/esnchez/weather_alert/domain/weather"
)

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

	for _, v := range mr.list {
		if v == wr {
			return fmt.Errorf("Weather registry already exists: %w", weather.ErrFailedToCreateRegistry)
		}
	}

	mr.list = append(mr.list, wr)

	return nil
}
