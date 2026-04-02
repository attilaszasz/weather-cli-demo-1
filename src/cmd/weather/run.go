package main

import (
	"context"
	"flag"
	"fmt"
	"io"

	"weather-cli-demo-1/src/internal/provider"
	"weather-cli-demo-1/src/internal/provider/openmeteo"
	"weather-cli-demo-1/src/internal/service"
	"weather-cli-demo-1/src/internal/validation"
)

type weatherGetter interface {
	GetCurrentWeather(ctx context.Context, latitude float64, longitude float64) (provider.WeatherData, error)
}

func run(args []string, stdout io.Writer, stderr io.Writer) error {
	defaultService := service.NewWeatherService(openmeteo.NewClient(nil))
	return runWithService(args, stdout, stderr, defaultService)
}

func runWithService(args []string, stdout io.Writer, stderr io.Writer, weatherService weatherGetter) error {
	flagSet := flag.NewFlagSet("weather", flag.ContinueOnError)
	flagSet.SetOutput(stderr)

	var latitude float64
	var longitude float64
	var latitudeSet bool
	var longitudeSet bool
	var help bool

	flagSet.Func("latitude", "latitude", func(value string) error {
		latitudeSet = true
		_, err := fmt.Sscan(value, &latitude)
		return err
	})
	flagSet.Func("longitude", "longitude", func(value string) error {
		longitudeSet = true
		_, err := fmt.Sscan(value, &longitude)
		return err
	})
	flagSet.BoolVar(&help, "help", false, "show help")
	flagSet.Usage = func() {
		_, _ = io.WriteString(stderr, usageText)
	}

	if err := flagSet.Parse(args); err != nil {
		return err
	}

	if help {
		_, _ = io.WriteString(stdout, usageText)
		return nil
	}

	if err := validation.ValidateArgs(flagSet.Args()); err != nil {
		return err
	}
	if err := validation.ValidateCoordinates(latitude, longitude, latitudeSet, longitudeSet); err != nil {
		return err
	}

	weather, err := weatherService.GetCurrentWeather(context.Background(), latitude, longitude)
	if err != nil {
		return err
	}

	_, err = fmt.Fprintf(stdout, "%+v\n", weather)
	return err
}
