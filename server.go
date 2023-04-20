package main

import (
	"fmt"
	"net/http"

	service "github.com/esnchez/weather_alert/service/weather"
	"github.com/gorilla/mux"
)

type Server struct {
	listenAddr string
	svc        *service.WeatherService
}

func NewServer(listenAddress string, svc *service.WeatherService) *Server {
	return &Server{
		listenAddr: listenAddress,
		svc:        svc,
	}
}

func (s *Server) Run() error {
	router := mux.NewRouter()

	router.HandleFunc("/weather/{cityname}", s.handleCreateRegistry).Methods("POST")
	router.HandleFunc("/weather", s.handleGetRegistry).Methods("GET")

	fmt.Println("Server running on port", s.listenAddr)
	return http.ListenAndServe(s.listenAddr, router)
}
