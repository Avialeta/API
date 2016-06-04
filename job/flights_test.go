package job_test

import (
	"fmt"
	"net/url"
	"testing"
	"time"

	"github.com/avialeta/api/cache"
	"github.com/avialeta/api/ovio"

	. "github.com/avialeta/api/job"
)

func init() {
	cache.LoadLocations()
}

func TestFetchFlights(t *testing.T) {
	outboundDate := fmt.Sprintf("%d-%02d-%02d", time.Now().Year(), time.Now().Month(), time.Now().Day())

	v := url.Values{
		"pointA":       {"MSQ"},
		"pointB":       {"LED"},
		"outboundDate": {outboundDate},
	}

	_, err := FetchFlights(v)
	if err != nil {
		t.Error(err)
	}
}

func TestFilterFlights(t *testing.T) {
	variants := ovio.Variants{
		"foo": ovio.Variant{
			Price:    30.00,
			Currency: "EUR",
			Segments: []ovio.Segment{{
				Flights: []ovio.Flight{},
			}},
		},
		"bar": ovio.Variant{
			Price:    20.00,
			Currency: "EUR",
			Segments: []ovio.Segment{{
				Flights: []ovio.Flight{},
			}},
		},
		"baz": ovio.Variant{
			Price:    10.00,
			Currency: "EUR",
			Segments: []ovio.Segment{{
				Flights: []ovio.Flight{},
			}},
		},
	}

	_variants := FilterVairants(variants)

	t.Log(_variants)
}
