package job_test

import (
	"testing"

	. "github.com/avialeta/api/job"
)

func TestFetchCities(t *testing.T) {
	FetchCities(RandomCountryCode)
}

func BenchmarkFetchCities(b *testing.B) {
	for cityCode := range Countries {
		FetchCities(cityCode)
	}
}
