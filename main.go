package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/esnchez/weather_alert/domain/weather"
)

var (
	urlBcn = "https://api.openweathermap.org/data/2.5/weather?q=barcelona&units=metric&appid=d6c6d8d30c6b59e827fb054180f82198"
)

func main() {

	resp, err := http.DefaultClient.Get(urlBcn)
	if err != nil {
		fmt.Printf("An error occurred fetching data %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Expected error code 200, got %v", resp.StatusCode)
	}

	defer resp.Body.Close()

	weatherResponse := &weather.WeatherJSONResp{}
	json.NewDecoder(resp.Body).Decode(weatherResponse)

	fmt.Println(weatherResponse )

	fmt.Println("Listening on port :8080")
	http.ListenAndServe(":8080", nil)
}
