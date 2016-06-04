package iatacodes_test

import (
	"testing"

	. "github.com/avialeta/api/iatacodes"
)

func TestFetchAirlines(t *testing.T) {
	if _, err := FetchAirlines(); err != nil {
		t.Error(err)
	}
}
