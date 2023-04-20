package main

import (
	"log"

	"github.com/esnchez/weather_alert/domain/weather/mysql"
	"github.com/esnchez/weather_alert/service/notification"
	service "github.com/esnchez/weather_alert/service/weather"
)

var (
	weatherUrl = "https://api.openweathermap.org/data/2.5/weather?q=cityname&units=metric&appid=d6c6d8d30c6b59e827fb054180f82198"
)

func main() {

	db, err := mysql.NewMySQLDatabase()
	if err != nil {
		log.Fatalf("DB connection error: %s", err)
	}

	cfg := notification.NewTwilioConfig()

	weatherSvc, err := service.NewWeatherService(
		service.WithMySQLWeatherRepository(db.GetDB()),
		service.WithTwilioService(cfg),
	)
	if err != nil {
		log.Fatalf("Error creating the weather service: %s", err)
	}

	server := NewServer(":5000", weatherSvc)
	log.Fatal(server.Run())
}
