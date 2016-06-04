package job

import (
	"github.com/avialeta/api/cache"
	"github.com/avialeta/api/log"
	"github.com/avialeta/api/ovio"
	"github.com/avialeta/api/search"
)

var Cities = make(map[string]City)

type City struct {
	Id          string
	Name        string
	CountryCode string
}

func (c City) Map() map[string]string {
	return map[string]string{
		"Id":          c.Id,
		"Name":        c.Name,
		"CountryCode": c.CountryCode,
	}
}

func RestoreCities() bool {
	cities := cache.LoadCities()
	if len(cities) == 0 {
		return false
	}

	for _, city := range cities {
		Cities[city["Id"]] = City{
			Id:          city["Id"],
			Name:        city["Name"],
			CountryCode: city["CountryCode"],
		}
	}

	return true
}

func fetchCities(chCountries <-chan Country) <-chan ovio.City {
	chCities := make(chan ovio.City)

	for i := 0; i < 10; i++ {
		go func() {
			for country := range chCountries {
				cities, err := ovio.FetchCities(country.Id)
				if err != nil {
					log.Error.Println(err)
				}

				for _, city := range cities {
					chCities <- city
				}
			}
		}()
	}

	return chCities
}

func processCities(chRawCities <-chan ovio.City) {
	go func() {
		for rawCity := range chRawCities {
			city := City{
				Id:          rawCity.Code,
				Name:        rawCity.Name,
				CountryCode: rawCity.Country,
			}

			Cities[city.Id] = city

			search.AddCity(mapCity(city))
			cache.SaveCities(city.Map())
		}

		search.CreateAirportsStr()
	}()
}
