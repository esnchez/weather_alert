package weather

type Repository interface {
	Save(*Weather) error
}