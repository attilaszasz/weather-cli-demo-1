package main

import (
	"bytes"
	"context"
	"errors"
	"strings"
	"testing"

	"weather-cli-demo-1/src/internal/provider"
)

type stubWeatherService struct {
	weather provider.WeatherData
	err     error
}

func (s stubWeatherService) GetCurrentWeather(_ context.Context, latitude float64, longitude float64) (provider.WeatherData, error) {
	return s.weather, s.err
}

func runCommand(t *testing.T, args []string, weatherService weatherGetter) (*bytes.Buffer, *bytes.Buffer, error) {
	t.Helper()
	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}
	err := runWithService(args, stdout, stderr, weatherService)
	return stdout, stderr, err
}

func TestRunHelp(t *testing.T) {
	stdout, _, err := runCommand(t, []string{"--help"}, stubWeatherService{})
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}
	if !strings.Contains(stdout.String(), "--latitude") {
		t.Fatalf("expected help output, got %q", stdout.String())
	}
}

func TestRunValidationFailures(t *testing.T) {
	tests := []struct {
		name string
		args []string
	}{
		{name: "missing longitude", args: []string{"--latitude", "10"}},
		{name: "missing latitude", args: []string{"--longitude", "20"}},
		{name: "latitude out of range", args: []string{"--latitude", "100", "--longitude", "20"}},
		{name: "longitude out of range", args: []string{"--latitude", "10", "--longitude", "200"}},
		{name: "extra arg", args: []string{"--latitude", "10", "--longitude", "20", "extra"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, _, err := runCommand(t, tt.args, stubWeatherService{})
			if err == nil {
				t.Fatal("expected validation error")
			}
		})
	}
}

func TestRunSuccess(t *testing.T) {
	stdout, _, err := runCommand(t, []string{"--latitude", "10", "--longitude", "20"}, stubWeatherService{
		weather: provider.WeatherData{Temperature: 20.5, WindSpeed: 4.1, WeatherCode: 2, ObservationTime: "2026-04-02T10:00"},
	})
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}
	if !strings.Contains(stdout.String(), "Temperature:20.5") {
		t.Fatalf("unexpected stdout: %q", stdout.String())
	}
}

func TestRunReturnsProviderError(t *testing.T) {
	_, _, err := runCommand(t, []string{"--latitude", "10", "--longitude", "20"}, stubWeatherService{err: errors.New("provider failed")})
	if err == nil {
		t.Fatal("expected provider error")
	}
}
