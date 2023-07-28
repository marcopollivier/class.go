package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type WeatherData struct {
	Country string `json:"country"`
	Date    string `json:"date"`
	Text    string `json:text`
}

func main() {

	var apiUrl = "http://apiadvisor.climatempo.com.br/api/v1/anl/synoptic/locale/BR?token=XXXX"

	resp, _ := http.Get(apiUrl)
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var weatherData []WeatherData

	json.Unmarshal(body, &weatherData)

	fmt.Println(weatherData[0].Country)
	fmt.Println(weatherData[0].Date)
	fmt.Println(weatherData[0].Text)
}
