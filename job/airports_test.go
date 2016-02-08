package job_test

import (
	"testing"

	. "github.com/avialeta/api/job"
)

func TestFetchAirports(t *testing.T) {
	FetchAirports(RandomCountryCode)
}
