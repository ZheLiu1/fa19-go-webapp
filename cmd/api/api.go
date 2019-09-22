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
	fmt.Println(vars["cityname"])
	resp, err := http.Get("http://api.openweathermap.org/data/2.5/weather?appid=470df66a992b911da9bfe98606f8461c&q=" + vars["cityname"])
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(body))
	w.Write(body)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/current/{cityname}", GetWeather)
	log.Fatal(http.ListenAndServe(":8080", router))
}
