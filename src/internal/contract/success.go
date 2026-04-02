package contract

import (
	"encoding/json"
	"io"

	"weather-cli-demo-1/src/internal/provider"
)

type SuccessResponse struct {
	Temperature     float64 `json:"temperature"`
	WindSpeed       float64 `json:"windSpeed"`
	WeatherCode     int     `json:"weatherCode"`
	ObservationTime string  `json:"observationTime"`
}

func NewSuccessResponse(weather provider.WeatherData) SuccessResponse {
	return SuccessResponse{
		Temperature:     weather.Temperature,
		WindSpeed:       weather.WindSpeed,
		WeatherCode:     weather.WeatherCode,
		ObservationTime: weather.ObservationTime,
	}
}

func WriteSuccessResponse(writer io.Writer, weather provider.WeatherData) error {
	encoder := json.NewEncoder(writer)
	return encoder.Encode(NewSuccessResponse(weather))
}
