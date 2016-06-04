package googleapis_test

import (
	"testing"

	. "github.com/avialeta/api/data/googleapis"
)

func TestFetchTimeZone(t *testing.T) {
	location := Location{53.8825, 28.0325}
	FetchTimeZone(location)
}
