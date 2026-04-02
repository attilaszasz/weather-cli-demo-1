package service

import (
	"context"
	"errors"
	"testing"

	"weather-cli-demo-1/src/internal/provider"
)

type stubProvider struct {
	request provider.Request
	result  provider.WeatherData
	err     error
}

func (s *stubProvider) CurrentWeather(_ context.Context, request provider.Request) (provider.WeatherData, error) {
	s.request = request
	return s.result, s.err
}

func TestGetCurrentWeatherDelegatesToProvider(t *testing.T) {
	stub := &stubProvider{result: provider.WeatherData{Temperature: 20, WindSpeed: 5, WeatherCode: 1, ObservationTime: "2026-04-02T10:00"}}
	service := NewWeatherService(stub)

	weather, err := service.GetCurrentWeather(context.Background(), 10, 20)
	if err != nil {
		t.Fatalf("GetCurrentWeather returned error: %v", err)
	}
	if stub.request.Latitude != 10 || stub.request.Longitude != 20 {
		t.Fatalf("unexpected request: %+v", stub.request)
	}
	if weather.ObservationTime != "2026-04-02T10:00" {
		t.Fatalf("unexpected weather: %+v", weather)
	}
}

func TestGetCurrentWeatherReturnsProviderError(t *testing.T) {
	stub := &stubProvider{err: errors.New("provider failed")}
	service := NewWeatherService(stub)

	_, err := service.GetCurrentWeather(context.Background(), 10, 20)
	if err == nil {
		t.Fatal("expected provider error")
	}
}
