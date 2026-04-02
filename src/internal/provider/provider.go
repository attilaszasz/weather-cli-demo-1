package provider

import "context"

type Request struct {
	Latitude  float64
	Longitude float64
}

type WeatherData struct {
	Temperature     float64 `json:"temperature"`
	WindSpeed       float64 `json:"windSpeed"`
	WeatherCode     int     `json:"weatherCode"`
	ObservationTime string  `json:"observationTime"`
}

type WeatherProvider interface {
	CurrentWeather(ctx context.Context, request Request) (WeatherData, error)
}
