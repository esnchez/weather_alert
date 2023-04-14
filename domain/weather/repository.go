package weather

type Repository interface {
	Save(*WeatherRegistry) error
}