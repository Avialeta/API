package search

import (
	"fmt"
	"strings"
)

var (
	CitiesCodes = make(map[string]*Location)
	CitiesNames = make(map[string]*Location)
	cities      string
)

func AddCity(city *Location) {
	CitiesCodes[city.Code] = city
	CitiesNames[city.Name] = city
}

func CreateCitiesStr() {
	cities = ""
	for city := range CitiesNames {
		cities += fmt.Sprintf("%s%s", city, SEPARATOR)
	}
	strings.Trim(cities, SEPARATOR)
}

func FindCities(pat string) []Location {
	locations := []Location{}

	names := Locations(pat, cities)
	for _, name := range names {
		location := Location{CitiesNames[name].Code, name, ""}
		locations = append(locations, location)
	}

	return locations
}
