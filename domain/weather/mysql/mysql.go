package mysql

import (
	"database/sql"
	"github.com/esnchez/weather_alert/domain/weather"
	
)

type MySqlRepository struct {
	db *sql.DB
}

func NewMySQLRepository(db *sql.DB) *MySqlRepository {
	return &MySqlRepository{
		db: db,
	}
}

func (r *MySqlRepository) GetAll() ([]*weather.WeatherRegistry, error) {

	query := `select * from registries`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	regs := []*weather.WeatherRegistry{}
	for rows.Next() {
		r := new(weather.WeatherRegistry)
		err := rows.Scan(
			&r.Id,
			&r.CityName,
			&r.Temperature,
			&r.StateCode,
			&r.StateDesc,
		)
		if err != nil {
			return nil, err
		}
		regs = append(regs, r)
	}

	return regs, nil
}

func (r *MySqlRepository) Save(wr *weather.WeatherRegistry) error {

	query := `insert into registries (
		id, cityname, temperature,  statecode, description) 
		values (?, ?, ?, ?, ?)`

	_, err := r.db.Query(query, wr.Id, wr.CityName, wr.Temperature, wr.StateCode, wr.StateDesc)
	if err != nil {
		return err
	}

	return nil
}

