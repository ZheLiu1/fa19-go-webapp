package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func GetWeather(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cityName := vars["cityName"]

	resp, err := http.Get(fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?appid=470df66a992b911da9bfe98606f8461c&q=%s", cityName))
	if err != nil {
		log.Printf("Error occurred during getting weather: %v", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error occurred during getting weather: %v", err)
	}
	log.Printf("Response for city %v is: %v", cityName, string(body))

	_, err = w.Write(body)
	if err != nil {
		log.Println("Error occurred during writing response.")
	}
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/current/{cityName}", GetWeather)

	host, port := "0.0.0.0", "8080"
	log.Printf("Server started on %v:%v", host, port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%v:%v", host, port), router))
}
