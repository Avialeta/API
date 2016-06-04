package search_test

import (
	"testing"

	. "github.com/avialeta/api/search"
)

func TestLocations(t *testing.T) {
	testTable := []struct {
		Pattern   string
		Locations string
		Expect    []string
	}{
		{"Foo", "Foo\tBar\tBaz\tBazz\tBazzz\t", []string{"Foo"}},
		{"Baz", "Foo\tBar\tBaz\tBazz\tBazzz\t", []string{"Baz", "Bazz", "Bazzz"}},
		{"Bazzz", "Foo\tBar\tBaz\tBazz\tBazzz\t", []string{"Bazzz"}},
	}

	for _, v := range testTable {
		sugg := Locations(v.Pattern, v.Locations)
		if len(sugg) == 0 {
			t.Error(v.Pattern)
		}
		t.Log(sugg)
	}
}
