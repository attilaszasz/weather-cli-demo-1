package service

import (
	"context"

	"weather-cli-demo-1/src/internal/provider"
)

type WeatherService struct {
	provider provider.WeatherProvider
}

func NewWeatherService(weatherProvider provider.WeatherProvider) *WeatherService {
	return &WeatherService{provider: weatherProvider}
}

func (s *WeatherService) GetCurrentWeather(ctx context.Context, latitude float64, longitude float64) (provider.WeatherData, error) {
	return s.provider.CurrentWeather(ctx, provider.Request{
		Latitude:  latitude,
		Longitude: longitude,
	})
}
