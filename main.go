package main

import (
	"fmt"
	"net/http"
)

var (
	urlBcn = "https://api.openweathermap.org/data/2.5/weather?q=barcelona&units=metric&appid=d6c6d8d30c6b59e827fb054180f82198"
)

func main() {

	resp, err := http.Get(urlBcn)
	if err != nil {
		fmt.Printf("An error occurred fetching data %v", err)
	}

	fmt.Println(resp.Body)

	fmt.Println("Listening on port :8080")
	http.ListenAndServe(":8080", nil)
}
