package googleapis_test

import (
	"testing"

	. "github.com/avialeta/api/data/googleapis"
)

func TestFetchLocation(t *testing.T) {
	_, err := FetchLocation("MSQ")
	if err != nil {
		t.Error(err)
	}
}
