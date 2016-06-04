package ovio_test

import (
	"fmt"
	"net/url"
	"testing"
	"time"

	. "github.com/avialeta/api/ovio"
)

func TestFetchFlights(t *testing.T) {
	values := url.Values{
		"pointA":       {"MSQ"},
		"pointB":       {"LIS"},
		"outboundDate": {fmt.Sprintf("%d-%02d-%02d", time.Now().Year(), time.Now().Month(), time.Now().Day())},
	}

	_, err := FetchFlights(values)
	if err != nil {
		t.Error(err)
	}
}
