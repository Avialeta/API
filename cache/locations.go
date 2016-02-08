package cache

import (
	"github.com/avialeta/api/log"
	"github.com/avialeta/api/search"
)

func SaveLocations() {
	log.Info.Print("SaveLocations --start--")
	defer log.Info.Print("SaveLocations --end--")

	client.Del("countries", "cities", "airports")

	for _, country := range search.CountriesCodes {
		client.LPush("countries", country.Code)
		client.HSet("countriesNames", country.Code, country.Name)
	}

	for _, city := range search.CitiesCodes {
		client.LPush("cities", city.Code)
		client.HSet("citiesNames", city.Code, city.Name)
	}

	for _, airport := range search.AirportsCodes {
		client.LPush("airports", airport.Code)
		client.HSet("airportsNames", airport.Code, airport.Name)
	}
}

func LoadLocations() bool {
	log.Info.Print("LoadLocations --start--")
	defer log.Info.Print("LoadLocations --end--")

	types := []string{"countries", "cities", "airports"}
	for _, t := range types {
		l, err := client.LLen(t)
		if err != nil {
			log.Error.Print(err)
			return false
		}

		if l == 0 {
			return false
		}
	}

	var codes []string
	var err error

	codes, err = client.LRange("countries", 0, -1)
	if err != nil {
		log.Error.Print(err)
	} else {
		countries, err := client.HGetAll("countriesNames")
		if err != nil {
			log.Error.Print(err)
		}

		for _, code := range codes {
			name, ok := countries[code]
			if ok {
				country := &search.Location{
					Code: code,
					Name: name,
				}
				search.AddCountry(country)
			}
		}
	}

	codes, err = client.LRange("cities", 0, -1)
	if err != nil {
		log.Error.Print(err)
	} else {
		cities, err := client.HGetAll("citiesNames")
		if err != nil {
			log.Error.Print(err)
		}

		for _, code := range codes {
			name, ok := cities[code]
			if ok {
				city := &search.Location{
					Code: code,
					Name: name,
				}
				search.AddCity(city)
			}
		}
	}

	codes, err = client.LRange("airports", 0, -1)
	if err != nil {
		log.Error.Print(err)
	} else {
		airports, err := client.HGetAll("airportsNames")
		if err != nil {
			log.Error.Print(err)
		}

		for _, code := range codes {
			name, ok := airports[code]
			if ok {
				airport := &search.Location{
					Code: code,
					Name: name,
				}
				search.AddAirport(airport)
			}
		}
	}

	return true
}
