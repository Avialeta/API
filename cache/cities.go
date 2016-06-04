package cache

import (
	"github.com/avialeta/api/log"
)

func RemoveCities() {
	codes, err := client.LRange("cities", 0, -1)
	if err != nil {
		log.Error.Println(err)
	} else {
		keys := make([]string, len(codes))

		for i, code := range codes {
			keys[i] = "city-" + code
		}

		keys = append(keys, "cities")

		client.Del(keys...)
	}
}

func SaveCities(city map[string]string) {
	if code, ok := city["Id"]; ok {
		client.LPush("cities", code)

		for k, v := range city {
			client.HSet("city-"+code, k, v)
		}
	}
}

func LoadCities() []map[string]string {
	cities := []map[string]string{}

	codes, err := client.LRange("cities", 0, -1)
	if err != nil {
		log.Error.Println(err)
		return cities
	} else {
		cities = make([]map[string]string, len(codes))
		for i, code := range codes {
			airline, err := client.HGetAll("city-" + code)
			if err != nil {
				log.Error.Println(err)
			} else {
				cities[i] = airline
			}
		}
	}

	return cities
}
