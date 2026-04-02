package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"strings"
	"testing"

	"weather-cli-demo-1/src/internal/contract"
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
		name     string
		args     []string
		category contract.ErrorCategory
		exitCode int
	}{
		{name: "missing longitude", args: []string{"--latitude", "10"}, category: contract.ErrorCategoryValidation, exitCode: contract.ExitCodeUsage},
		{name: "missing latitude", args: []string{"--longitude", "20"}, category: contract.ErrorCategoryValidation, exitCode: contract.ExitCodeUsage},
		{name: "latitude out of range", args: []string{"--latitude", "100", "--longitude", "20"}, category: contract.ErrorCategoryValidation, exitCode: contract.ExitCodeUsage},
		{name: "longitude out of range", args: []string{"--latitude", "10", "--longitude", "200"}, category: contract.ErrorCategoryValidation, exitCode: contract.ExitCodeUsage},
		{name: "extra arg", args: []string{"--latitude", "10", "--longitude", "20", "extra"}, category: contract.ErrorCategoryValidation, exitCode: contract.ExitCodeUsage},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stdout, stderr, err := runCommand(t, tt.args, stubWeatherService{})
			if err == nil {
				t.Fatal("expected validation error")
			}
			if stdout.Len() != 0 {
				t.Fatalf("expected empty stdout, got %q", stdout.String())
			}

			var payload contract.ErrorResponse
			if err := json.Unmarshal(stderr.Bytes(), &payload); err != nil {
				t.Fatalf("expected valid JSON stderr, got %q: %v", stderr.String(), err)
			}
			if payload.Category != tt.category {
				t.Fatalf("unexpected category: %#v", payload)
			}
			if payload.ExitCode != tt.exitCode {
				t.Fatalf("unexpected exit code: %#v", payload)
			}
		})
	}
}

func TestRunSuccess(t *testing.T) {
	stdout, stderr, err := runCommand(t, []string{"--latitude", "10", "--longitude", "20"}, stubWeatherService{
		weather: provider.WeatherData{Temperature: 20.5, WindSpeed: 4.1, WeatherCode: 2, ObservationTime: "2026-04-02T10:00"},
	})
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}
	if stderr.Len() != 0 {
		t.Fatalf("expected empty stderr, got %q", stderr.String())
	}

	var payload contract.SuccessResponse
	if err := json.Unmarshal(stdout.Bytes(), &payload); err != nil {
		t.Fatalf("expected valid JSON stdout, got %q: %v", stdout.String(), err)
	}

	expected := contract.SuccessResponse{
		Temperature:     20.5,
		WindSpeed:       4.1,
		WeatherCode:     2,
		ObservationTime: "2026-04-02T10:00",
	}
	if payload != expected {
		t.Fatalf("unexpected payload: %#v", payload)
	}
}

func TestRunReturnsProviderError(t *testing.T) {
	stdout, stderr, err := runCommand(t, []string{"--latitude", "10", "--longitude", "20"}, stubWeatherService{err: errors.New("provider failed")})
	if err == nil {
		t.Fatal("expected provider error")
	}
	if stdout.Len() != 0 {
		t.Fatalf("expected empty stdout, got %q", stdout.String())
	}

	var payload contract.ErrorResponse
	if err := json.Unmarshal(stderr.Bytes(), &payload); err != nil {
		t.Fatalf("expected valid JSON stderr, got %q: %v", stderr.String(), err)
	}
	if payload.Category != contract.ErrorCategoryInternal {
		t.Fatalf("unexpected payload: %#v", payload)
	}
	if payload.ExitCode != contract.ExitCodeInternal {
		t.Fatalf("unexpected payload: %#v", payload)
	}
}

func TestRunReturnsTransportProviderError(t *testing.T) {
	stdout, stderr, err := runCommand(t, []string{"--latitude", "10", "--longitude", "20"}, stubWeatherService{err: errors.New("execute open-meteo request: timeout")})
	if err == nil {
		t.Fatal("expected provider error")
	}
	if stdout.Len() != 0 {
		t.Fatalf("expected empty stdout, got %q", stdout.String())
	}

	var payload contract.ErrorResponse
	if err := json.Unmarshal(stderr.Bytes(), &payload); err != nil {
		t.Fatalf("expected valid JSON stderr, got %q: %v", stderr.String(), err)
	}
	if payload.Category != contract.ErrorCategoryTransport || payload.ExitCode != contract.ExitCodeDownstream {
		t.Fatalf("unexpected payload: %#v", payload)
	}
}
