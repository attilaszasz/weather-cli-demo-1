package openmeteo

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"weather-cli-demo-1/src/internal/provider"
)

func TestCurrentWeatherBuildsRequestAndParsesResponse(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("latitude") != "47.4979" {
			t.Fatalf("unexpected latitude: %s", r.URL.Query().Get("latitude"))
		}
		if r.URL.Query().Get("longitude") != "19.0402" {
			t.Fatalf("unexpected longitude: %s", r.URL.Query().Get("longitude"))
		}
		if !strings.Contains(r.URL.Query().Get("current"), "temperature_2m") {
			t.Fatalf("missing current query: %s", r.URL.Query().Get("current"))
		}
		_, _ = w.Write([]byte(`{"latitude":47.4979,"longitude":19.0402,"current":{"temperature_2m":21.5,"wind_speed_10m":4.2,"weather_code":3,"time":"2026-04-02T10:00"}}`))
	}))
	defer server.Close()

	client := NewClientWithEndpoint(&http.Client{Timeout: time.Second}, server.URL)
	weather, err := client.CurrentWeather(context.Background(), provider.Request{Latitude: 47.4979, Longitude: 19.0402})
	if err != nil {
		t.Fatalf("CurrentWeather returned error: %v", err)
	}

	if weather.Temperature != 21.5 || weather.WindSpeed != 4.2 || weather.WeatherCode != 3 || weather.ObservationTime != "2026-04-02T10:00" {
		t.Fatalf("unexpected weather data: %+v", weather)
	}
}

func TestCurrentWeatherRejectsMalformedPayload(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`{"current":{}}`))
	}))
	defer server.Close()

	client := NewClientWithEndpoint(&http.Client{Timeout: time.Second}, server.URL)
	_, err := client.CurrentWeather(context.Background(), provider.Request{Latitude: 1, Longitude: 1})
	if err == nil {
		t.Fatal("expected error for malformed payload")
	}
}

func TestCurrentWeatherRejectsNonOKStatus(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadGateway)
	}))
	defer server.Close()

	client := NewClientWithEndpoint(&http.Client{Timeout: time.Second}, server.URL)
	_, err := client.CurrentWeather(context.Background(), provider.Request{Latitude: 1, Longitude: 1})
	if err == nil {
		t.Fatal("expected error for non-200 status")
	}
}
