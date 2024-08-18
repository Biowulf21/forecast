package weather

import (
	"encoding/json"
	"errors"
	"io"
	"math"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

func FetchWeatherAPIKey() (string, error) {
	apiKey := os.Getenv("FORECAST_API")

	if apiKey == "" {
		return "", errors.New("No OpenWeatherMap.org API key found in environment variables.\n")
	}

	return apiKey, nil

}

type Weather struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Temperature float64 `json:"temp"`
	Humidity    int8    `json:"humidity"`
}

type WeatherResponse struct {
	Weather []struct {
		Description string `json:"description"`
		Name        string `json:"main"`
	} `json:"weather"`
	Main struct {
		Temperature float64 `json:"temp"`
		Humidity    int8    `json:"humidity"`
	} `json:"main"`
}

func FetchWeather(apiKey string, cityName string) (Weather, error) {
	position, err := fetchPosition(apiKey, cityName)

	if err != nil {
		return Weather{}, errors.New("Could not fetch Position")
	}

	long := strconv.FormatFloat(position.Latitude, 'f', -1, 32)
	lat := strconv.FormatFloat(position.Latitude, 'f', -1, 32)

	weatherUrl, err := url.Parse("https://api.openweathermap.org/data/2.5/weather")
	queryParams := weatherUrl.Query()

	queryParams.Add("lat", lat)
	queryParams.Add("lon", long)
	queryParams.Add("units", "metric")
	queryParams.Add("appid", apiKey)

	weatherUrl.RawQuery = queryParams.Encode()

	resp, responseError := http.Get(weatherUrl.String())

	if responseError != nil {
		return Weather{}, errors.New("Weather API unavailable.")
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return Weather{}, errors.New("Something went wrong.")
	}

	body, readErr := io.ReadAll(resp.Body)

	if readErr != nil {
		return Weather{}, errors.New("Something went wrong with reading the weather data.")
	}

	var weatherResponse WeatherResponse

	marshallErr := json.Unmarshal(body, &weatherResponse)

	if marshallErr != nil {
		return Weather{}, errors.New("Something went wrong with marshalling the data")
	}

	weather := Weather{
		Name:        weatherResponse.Weather[0].Name,
		Description: weatherResponse.Weather[0].Description,
		Temperature: weatherResponse.Main.Temperature,
		Humidity:    weatherResponse.Main.Humidity,
	}

	return weather, nil
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}

type Position struct {
	Longitude float64 `json:"lon"`
	Latitude  float64 `json:"lat"`
}

func fetchPosition(apiKey string, cityName string) (Position, error) {

	positionUrl, err := url.Parse("http://api.openweathermap.org/geo/1.0/direct")
	queryParams := positionUrl.Query()

	queryParams.Add("q", cityName)
	queryParams.Add("appid", apiKey)

	positionUrl.RawQuery = queryParams.Encode()

	resp, err := http.Get(positionUrl.String())
	if err != nil {
		return Position{}, errors.New("Cannot find your location")
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return Position{}, errors.New("Geocoding API not available")
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return Position{}, errors.New("Failed to read response body")
	}

	var positions []Position
	err = json.Unmarshal(body, &positions)
	if err != nil {
		return Position{}, errors.New("Failed to parse the API data")
	}

	if len(positions) == 0 {
		return Position{}, errors.New("No location data found")
	}

	return positions[0], nil
}
