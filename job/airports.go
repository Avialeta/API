package job

import (
	"github.com/avialeta/api/cache"
	"github.com/avialeta/api/log"
	"github.com/avialeta/api/ovio"
	"github.com/avialeta/api/search"
)

var Airports = make(map[string]Airport)

type Airport struct {
	Id          string
	Name        string
	CityCode    string
	CountryCode string
	TimeZone    string
}

func (a Airport) Map() map[string]string {
	return map[string]string{
		"Id":          a.Id,
		"Name":        a.Name,
		"CityCode":    a.CityCode,
		"CountryCode": a.CountryCode,
		"TimeZone":    a.TimeZone,
	}
}

func RestoreAirports() bool {
	airports := cache.LoadAirports()
	if len(airports) == 0 {
		return false
	}

	for _, airport := range airports {
		Airports[airport["Id"]] = Airport{
			Id:          airport["Id"],
			Name:        airport["Name"],
			CityCode:    airport["CityCode"],
			CountryCode: airport["CountryCode"],
			TimeZone:    airport["TimeZone"],
		}
	}

	return true
}

func fetchAirports(chCountries <-chan Country) chan ovio.Airport {
	chAirports := make(chan ovio.Airport)

	for i := 0; i < 10; i++ {
		go func() {
			for country := range chCountries {
				airports, err := ovio.FetchAirports(country.Id)
				if err != nil {
					log.Error.Println(err)
				}

				for i := range airports {
					chAirports <- airports[i]
				}
			}
		}()
	}

	return chAirports
}

func processAirports(chRawAirports <-chan ovio.Airport) {
	go func() {
		for rawAirport := range chRawAirports {
			airport := Airport{
				Id:          rawAirport.Id,
				Name:        rawAirport.Name,
				CityCode:    rawAirport.City,
				CountryCode: rawAirport.Country,
			}

			location := fetchLocations(Countries[airport.CountryCode].Name + ", " + Cities[airport.CityCode].Name)
			timeZone := fetchTimeZones(location)
			airport.TimeZone = timeZone

			Airports[airport.Id] = airport

			search.AddAirport(mapAirport(airport))
			cache.SaveAirports(airport.Map())
		}

		search.CreateCitiesStr()
	}()
}
