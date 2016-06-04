package job

import (
	"encoding/json"
	"net/url"

	//"github.com/avialeta/api/cache"
	//"github.com/avialeta/api/log"
	"github.com/avialeta/api/search"
)

/*func Locations() {
	if !cache.LoadLocations() {
		FetchLocations()
		cache.SaveLocations()
	}

	search.CreateCitiesStr()
	search.CreateAirportsStr()
}

func FetchLocations() {
	log.Info.Print("FetchLocations --start--")
	defer log.Info.Print("FetchLocations --end--")

	if err := FetchCountries(); err != nil {
		log.Error.Print(err)
		return
	}

	for id, country := range Countries {
		search.AddCountry(mapCountry(country))

		if err := FetchCities(id); err != nil {
			log.Error.Print(err)
		} else {
			for _, city := range Cities {
				search.AddCity(mapCity(city))
			}
		}

		if err := FetchAirports(id); err != nil {
			log.Error.Print(err)
		} else {
			for _, airport := range Airports {
				search.AddAirport(mapAirport(airport))
			}
		}
	}
}*/

func mapCountry(country Country) *search.Location {
	return &search.Location{
		Code: country.Id,
		Name: country.Name,
	}
}

func mapCity(country City) *search.Location {
	return &search.Location{
		Code: country.Id,
		Name: country.Name,
	}
}

func mapAirport(country Airport) *search.Location {
	return &search.Location{
		Code: country.Id,
		Name: country.Name,
		CityCode: country.CityCode,
	}
}

func SearchLocations(v url.Values) ([]byte, error) {
	pat, ok := v["search"]
	if !ok {
		return nil, nil
	}

	cities := search.FindCities(pat[0])
	airports := search.FindAirports(pat[0])
	locations := make([]search.Location, len(cities)+len(airports))
	copy(locations[:len(cities)], cities)
	copy(locations[len(airports):], airports)

	if len(locations) == 0 {
		return nil, nil
	}

	data, err := json.Marshal(locations)
	if err != nil {
		return nil, err
	}

	return data, nil
}
