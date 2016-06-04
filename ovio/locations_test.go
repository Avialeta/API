package ovio_test

import (
	"testing"

	. "github.com/avialeta/api/ovio"
)

var testCountries = []struct {
	Code string
}{{"BY"}, {"RU"}}

func TestFetchCountries(t *testing.T) {
	_, err := FetchCountries()
	if err != nil {
		t.Error(err)
	}
}

func TestFetchCities(t *testing.T) {
	for _, country := range testCountries {
		_, err := FetchCities(country.Code)
		if err != nil {
			t.Error(err)
		}
	}
}

func TestFetchAirports(t *testing.T) {
	for _, country := range testCountries {
		_, err := FetchAirports(country.Code)
		if err != nil {
			t.Error(err)
		}
	}
}
