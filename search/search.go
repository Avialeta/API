package search

import (
	"strings"

	"golang.org/x/text/language"
	"golang.org/x/text/search"
)

const (
	SEPARATOR = "\t"
)

var (
	options = []search.Option{
		search.IgnoreCase,
		search.IgnoreWidth,
	}

	matcher = search.New(language.English, options...)
)

type Location struct {
	Code string
	Name string
	CityCode string
}

type index struct {
	start int
	end   int
}

// Get cities suggestions
func Locations(pat, l string) []string {
	if pat == "" {
		return make([]string, 0, 0)
	}

	indexes := []index{}
	suggestions := []string{}

	// Get matched indexes.
	var start, end, offset int
	for {
		start, end = matcher.IndexString(l[offset:], pat, 0)
		if start == -1 || end == -1 {
			break
		}

		indexes = append(indexes, index{offset + start, offset + end})
		offset += end
	}

	// Get suggested locations
	for _, index := range indexes {
		start := strings.LastIndex(l[:index.start], SEPARATOR)
		if start == -1 {
			start = 0
		} else {
			start += 1
		}

		end := strings.Index(l[index.end:], SEPARATOR)
		if end == -1 {
			end = len(l)
		} else {
			end = index.end + end
		}

		suggestions = append(suggestions, l[start:end])
	}

	return suggestions
}

func GetCodeByName(name string) string {
	code := name

	if airport, ok := AirportsNames[code]; ok {
		return airport.Code

	}

	//log.Debug.Println("AirportsNames", code)

	/*if airport, ok := AirportsCitiesNames[name]; ok {
		return airport.Code
	}*/

	if city, ok := CitiesNames[code]; ok {
		code = city.Code

		if num, ok := CityAirportsNum[code]; ok && num > 1 {
			// TODO: Make constant and write description.
			code += "c"
		}
	}

	/*log.Debug.Println("CitiesNames", code)
	log.Debug.Println("CityAirportsNum", CityAirportsNum)*/

	return code
}

func GetNameByCode(code string) string {
	if airport, ok := AirportsCodes[code]; ok {
		return airport.Name
	}

	if city, ok := CitiesCodes[code]; ok {
		return city.Name
	}

	return code
}
