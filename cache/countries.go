package cache

import (
	"github.com/avialeta/api/log"
)

func RemoveCountries() {
	codes, err := client.LRange("countries", 0, -1)
	if err != nil {
		log.Error.Println(err)
	} else {
		keys := make([]string, len(codes))

		for i, code := range codes {
			keys[i] = "country-" + code
		}

		keys = append(keys, "countries")

		client.Del(keys...)
	}
}

func SaveCountries(country map[string]string) {
	if code, ok := country["Id"]; ok {
		client.LPush("countries", code)

		for k, v := range country {
			client.HSet("country-"+code, k, v)
		}
	}
}

func LoadCountries() []map[string]string {
	countries := []map[string]string{}

	codes, err := client.LRange("countries", 0, -1)
	if err != nil {
		log.Error.Println(err)
		return countries
	} else {
		countries = make([]map[string]string, len(codes))
		for i, code := range codes {
			airline, err := client.HGetAll("country-" + code)
			if err != nil {
				log.Error.Println(err)
			} else {
				countries[i] = airline
			}
		}
	}

	return countries
}
