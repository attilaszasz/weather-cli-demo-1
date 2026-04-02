package openmeteo

import "net/url"

const baseURL = "https://api.open-meteo.com/v1/forecast"

func buildURL(endpoint string, latitude float64, longitude float64) (string, error) {
	parsed, err := url.Parse(endpoint)
	if err != nil {
		return "", err
	}

	query := parsed.Query()
	query.Set("latitude", formatCoordinate(latitude))
	query.Set("longitude", formatCoordinate(longitude))
	query.Set("current", "temperature_2m,wind_speed_10m,weather_code,time")
	parsed.RawQuery = query.Encode()

	return parsed.String(), nil
}
