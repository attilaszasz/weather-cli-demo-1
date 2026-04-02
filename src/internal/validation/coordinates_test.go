package validation

import "testing"

func TestValidateCoordinates(t *testing.T) {
	tests := []struct {
		name         string
		latitude     float64
		longitude    float64
		latitudeSet  bool
		longitudeSet bool
		wantErr      bool
	}{
		{name: "valid", latitude: 10, longitude: 20, latitudeSet: true, longitudeSet: true},
		{name: "missing latitude", longitude: 20, longitudeSet: true, wantErr: true},
		{name: "missing longitude", latitude: 10, latitudeSet: true, wantErr: true},
		{name: "latitude out of range", latitude: 100, longitude: 20, latitudeSet: true, longitudeSet: true, wantErr: true},
		{name: "longitude out of range", latitude: 10, longitude: 200, latitudeSet: true, longitudeSet: true, wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateCoordinates(tt.latitude, tt.longitude, tt.latitudeSet, tt.longitudeSet)
			if (err != nil) != tt.wantErr {
				t.Fatalf("ValidateCoordinates() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateArgs(t *testing.T) {
	if err := ValidateArgs(nil); err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if err := ValidateArgs([]string{"extra"}); err == nil {
		t.Fatal("expected error for extra args")
	}
}
