package validation

import "fmt"

func ValidateCoordinates(latitude float64, longitude float64, latitudeSet bool, longitudeSet bool) error {
	if !latitudeSet {
		return fmt.Errorf("latitude is required")
	}
	if !longitudeSet {
		return fmt.Errorf("longitude is required")
	}
	if latitude < -90 || latitude > 90 {
		return fmt.Errorf("latitude must be between -90 and 90")
	}
	if longitude < -180 || longitude > 180 {
		return fmt.Errorf("longitude must be between -180 and 180")
	}
	return nil
}
