package job

import (
	"github.com/avialeta/api/log"
	"github.com/avialeta/api/search"
)

func Scheduler() {
	log.Info.Print("Run Scheduler")

	if RestoreCountries() {
		RestoreCities()
		RestoreAirports()
		RestoreAirlines()

		for id := range Countries {
			search.AddCountry(mapCountry(Countries[id]))
		}

		for id := range Cities {
			search.AddCity(mapCity(Cities[id]))
		}

		for id := range Airports {
			search.AddAirport(mapAirport(Airports[id]))
		}

		search.CreateCitiesStr()
		search.CreateAirportsStr()

		log.Info.Println("Restore")
		return
	}

	chRawCountries := fetchCountries()

	chCountriesCities, chCountriesAirports := processCountries(chRawCountries)

	rawCities := fetchCities(chCountriesCities)
	rawAirports := fetchAirports(chCountriesAirports)

	processCities(rawCities)
	processAirports(rawAirports)

	go fetchAirlines()
}
