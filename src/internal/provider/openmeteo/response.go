package openmeteo

type responseEnvelope struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Current   struct {
		Temperature2M float64 `json:"temperature_2m"`
		WindSpeed10M  float64 `json:"wind_speed_10m"`
		WeatherCode   int     `json:"weather_code"`
		Time          string  `json:"time"`
	} `json:"current"`
}
