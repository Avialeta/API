package job

import (
	"github.com/avialeta/api/cache"
	"github.com/avialeta/api/log"
	"github.com/avialeta/api/ovio"
	"github.com/avialeta/api/search"
)

var Countries = make(map[string]Country)

type Country struct {
	Id   string
	Name string
}

func (c Country) Map() map[string]string {
	return map[string]string{
		"Id":   c.Id,
		"Name": c.Name,
	}
}

func RestoreCountries() bool {
	countries := cache.LoadCountries()
	if len(countries) == 0 {
		return false
	}

	for _, country := range countries {
		Countries[country["Id"]] = Country{
			Id:   country["Id"],
			Name: country["Name"],
		}
	}

	return true
}

func fetchCountries() <-chan ovio.Country {
	chRawCountries := make(chan ovio.Country)

	go func() {
		countries, err := ovio.FetchCountries()
		if err != nil {
			log.Error.Println(err)
		}

		for _, country := range countries {
			chRawCountries <- country
		}
		close(chRawCountries)
	}()

	return chRawCountries
}

func processCountries(rawCountries <-chan ovio.Country) (<-chan Country, <-chan Country) {
	chCountriesCities := make(chan Country)
	chCountriesAirports := make(chan Country)

	go func() {
		for rawCountry := range rawCountries {
			if rawCountry.Continent != "europe" {
				continue
			}

			country := Country{
				Id:   rawCountry.Id,
				Name: rawCountry.Name,
			}

			chCountriesCities <- country
			chCountriesAirports <- country

			Countries[country.Id] = country

			search.AddCountry(mapCountry(country))
			cache.SaveCountries(country.Map())
		}
		close(chCountriesCities)
		close(chCountriesAirports)
	}()

	return chCountriesCities, chCountriesAirports
}
