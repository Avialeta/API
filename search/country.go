package search

var (
	CountriesCodes = make(map[string]*Location)
	CountriesNames = make(map[string]*Location)
)

func AddCountry(country *Location) {
	CountriesCodes[country.Code] = country
	CountriesNames[country.Name] = country
}
