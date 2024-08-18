package main

import (
	"fmt"
	"forecast/weather"
	"os"
)

func main() {
	apiKey, err := weather.FetchWeatherAPIKey()
	if err != nil {
		fmt.Println(err)
	}

	var city string
	args := os.Args

	if len(args) > 1 {
		city = args[1]
	} else {
		city = "Cagayan de Oro City"
	}

	weather, weatherErr := weather.FetchWeather(apiKey, city)

	if weatherErr != nil {
		fmt.Println(weatherErr)
	}

	result := fmt.Sprintf("Weather for %s is %s (%s), which has a humidity of %d%%, and a temperature of %.2f°C", city, weather.Name, weather.Description, weather.Humidity, weather.Temperature)
	fmt.Println(result)
}
