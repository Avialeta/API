package job

import (
	"github.com/avialeta/api/cache"
	"github.com/avialeta/api/iatacodes"
	"github.com/avialeta/api/log"
)

var Airlines map[string]Airline = make(map[string]Airline)

type Airline struct {
	Icao string
	Code string
	Name string
}

func (a Airline) Map() map[string]string {
	return map[string]string{
		"Icao": a.Icao,
		"Code": a.Code,
		"Name": a.Name,
	}
}

func RestoreAirlines() bool {
	airlines := cache.LoadAirlines()
	if len(airlines) == 0 {
		return false
	}

	for _, airline := range airlines {
		Airlines[airline["Code"]] = Airline{
			Code: airline["Code"],
			Icao: airline["Icao"],
			Name: airline["Name"],
		}
	}

	return true
}

func fetchAirlines() {
	airlines, err := iatacodes.FetchAirlines()
	if err != nil {
		log.Error.Println(err)
	}

	for i := range airlines {
		airline := Airline{
			Code: airlines[i].Code,
			Icao: airlines[i].Icao,
			Name: airlines[i].Name,
		}

		Airlines[airlines[i].Code] = airline
		cache.SaveAirline(airline.Map())
	}
}
