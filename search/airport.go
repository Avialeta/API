package search

import (
	"fmt"
	"strings"
)

var (
	AirportsCodes = make(map[string]*Location)
	AirportsNames = make(map[string]*Location)
	CityAirportsNum = make(map[string]int)
	airports      string
)

func AddAirport(airport *Location) {
	AirportsCodes[airport.Code] = airport
	AirportsNames[airport.Name] = airport
	CityAirportsNum[airport.CityCode] = CityAirportsNum[airport.CityCode] + 1
}

func CreateAirportsStr() {
	airports = ""
	for airport := range AirportsNames {
		airports += fmt.Sprintf("%s%s", airport, SEPARATOR)
	}
	strings.Trim(airports, SEPARATOR)
}

func FindAirports(pat string) []Location {
	locations := []Location{}

	names := Locations(pat, airports)
	for _, name := range names {
		location := Location{AirportsNames[name].Code, name, ""}
		locations = append(locations, location)
	}

	return locations
}
