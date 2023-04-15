package weather

type Repository interface {
	GetAll() ([]*WeatherRegistry, error)
	Save(*WeatherRegistry) error
}
