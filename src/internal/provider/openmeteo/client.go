package openmeteo

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"weather-cli-demo-1/src/internal/provider"
)

const defaultTimeout = 3 * time.Second

type Client struct {
	endpoint   string
	httpClient *http.Client
}

func NewClient(httpClient *http.Client) *Client {
	client := httpClient
	if client == nil {
		client = &http.Client{Timeout: defaultTimeout}
	}

	if client.Timeout == 0 {
		client.Timeout = defaultTimeout
	}

	return &Client{
		endpoint:   baseURL,
		httpClient: client,
	}
}

func NewClientWithEndpoint(httpClient *http.Client, endpoint string) *Client {
	client := NewClient(httpClient)
	if endpoint != "" {
		client.endpoint = endpoint
	}
	return client
}

func (c *Client) CurrentWeather(ctx context.Context, request provider.Request) (provider.WeatherData, error) {
	requestURL, err := buildURL(c.endpoint, request.Latitude, request.Longitude)
	if err != nil {
		return provider.WeatherData{}, fmt.Errorf("build open-meteo request: %w", err)
	}

	httpRequest, err := http.NewRequestWithContext(ctx, http.MethodGet, requestURL, nil)
	if err != nil {
		return provider.WeatherData{}, fmt.Errorf("create open-meteo request: %w", err)
	}

	response, err := c.httpClient.Do(httpRequest)
	if err != nil {
		return provider.WeatherData{}, fmt.Errorf("execute open-meteo request: %w", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return provider.WeatherData{}, fmt.Errorf("open-meteo returned status %d", response.StatusCode)
	}

	var payload responseEnvelope
	if err := json.NewDecoder(response.Body).Decode(&payload); err != nil {
		return provider.WeatherData{}, fmt.Errorf("decode open-meteo response: %w", err)
	}

	if payload.Current.Time == "" {
		return provider.WeatherData{}, errors.New("open-meteo response missing current time")
	}

	return provider.WeatherData{
		Temperature:     payload.Current.Temperature2M,
		WindSpeed:       payload.Current.WindSpeed10M,
		WeatherCode:     payload.Current.WeatherCode,
		ObservationTime: payload.Current.Time,
	}, nil
}

func formatCoordinate(value float64) string {
	return strconv.FormatFloat(value, 'f', -1, 64)
}
