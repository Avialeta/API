package job

import (
	"math/rand"
	"testing"
)

var RandomCountryCode string

func init() {
	FetchCountries()

	codes := []string{}
	for code := range Countries {
		codes = append(codes, code)
	}

	RandomCountryCode = codes[rand.Intn(len(codes))]
}

func TestFetchCountries(t *testing.T) {
	if err := FetchCountries(); err != nil {
		t.Error(err)
	}
}
