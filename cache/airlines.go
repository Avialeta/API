package cache

import (
	"github.com/avialeta/api/log"
)

func RemoveAirlines() {
	codes, err := client.LRange("airlines", 0, -1)
	if err != nil {
		log.Error.Print(err)
	} else {
		keys := make([]string, len(codes))

		for i, code := range codes {
			keys[i] = "airline-" + code
		}

		keys = append(keys, "airlines")

		client.Del(keys...)
	}
}

func SaveAirline(airline map[string]string) {
	if code, ok := airline["Code"]; ok {
		client.LPush("airlines", code)

		for k, v := range airline {
			client.HSet("airline-"+code, k, v)
		}
	}
}

func LoadAirlines() []map[string]string {
	airlines := []map[string]string{}

	codes, err := client.LRange("airlines", 0, -1)
	if err != nil {
		log.Error.Print(err)
		return airlines
	} else {
		airlines = make([]map[string]string, len(codes))
		for i, code := range codes {
			airline, err := client.HGetAll("airline-" + code)
			if err != nil {
				log.Error.Print(err)
			} else {
				airlines[i] = airline
			}
		}
	}

	return airlines
}
