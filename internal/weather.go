package internal

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
)

func GetWeather(locationName string) (Weather, error) {
	weather_url := "https://api.weatherapi.com/v1/current.json?key=" + os.Getenv("WEATHER_TOKEN") + "&lang=uk&q=" + locationName + "&aqi=no"

	response, err := http.Get(weather_url)
	if err != nil {
		return Weather{}, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return Weather{}, err
	}

	var weatherData Weather
	err = json.Unmarshal(body, &weatherData)
	if err != nil {
		return Weather{}, err
	}
	return weatherData, err
}
