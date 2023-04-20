package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"github.com/esnchez/weather_alert/domain/weather"
	"github.com/gorilla/mux"
	"errors"


)

type GenericError struct {
	Message string `json:"message"`
}

var (
	ErrCityNotFound = errors.New("This city is not found or not available to fetch its weather prevision")
)

func (s *Server) handleCreateRegistry(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	city := mux.Vars(r)["cityname"]
	url := strings.Replace(weatherUrl, "cityname", city, 1)

	jsonResp, err := getWeatherJson(url)
	if err != nil {
		log.Println(err)
		json.NewEncoder(w).Encode(&GenericError{Message: ErrCityNotFound.Error()})
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := s.svc.CreateRegistry(jsonResp); err != nil {
		json.NewEncoder(w).Encode(&GenericError{Message: err.Error()})
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Registry created successfully!"))
}

func (s *Server) handleGetRegistry(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	list, err := s.svc.GetAllRegistries()
	if err != nil {
		log.Printf("Error retrieving weather registries data: %v", err)
		json.NewEncoder(w).Encode(&GenericError{Message: err.Error()})
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(list)
}

func getWeatherJson(url string) (*weather.WeatherJSONResp, error) {

	resp, err := http.DefaultClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("An error occurred fetching data from Weather API %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Expected HTPP status code 200, got %v", resp.StatusCode)
	}
	defer resp.Body.Close()

	weaResp := &weather.WeatherJSONResp{}
	json.NewDecoder(resp.Body).Decode(weaResp)
	return weaResp, nil
}

