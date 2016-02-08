package cache

import (
	"github.com/avialeta/api/log"
)

func RemoveAirports() {
	codes, err := client.LRange("airports", 0, -1)
	if err != nil {
		log.Error.Println(err)
	} else {
		keys := make([]string, len(codes))

		for i, code := range codes {
			keys[i] = "airport-" + code
		}

		keys = append(keys, "airports")

		client.Del(keys...)
	}
}

func SaveAirports(airline map[string]string) {
	if code, ok := airline["Id"]; ok {
		client.LPush("airports", code)

		for k, v := range airline {
			client.HSet("airport-"+code, k, v)
		}
	}
}

func LoadAirports() []map[string]string {
	airports := []map[string]string{}

	codes, err := client.LRange("airports", 0, -1)
	if err != nil {
		log.Error.Println(err)
		return airports
	} else {
		airports = make([]map[string]string, len(codes))
		for i, code := range codes {
			airline, err := client.HGetAll("airport-" + code)
			if err != nil {
				log.Error.Println(err)
			} else {
				airports[i] = airline
			}
		}
	}

	return airports
}
