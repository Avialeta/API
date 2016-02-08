package job

import (
	"github.com/avialeta/api/data/googleapis"
	"github.com/avialeta/api/log"
)

func fetchLocations(address string) googleapis.Location {
	location, err := googleapis.FetchLocation(address)
	if err != nil {
		log.Error.Println(err)
	}

	return location
}

func fetchTimeZones(location googleapis.Location) string {
	timezone, err := googleapis.FetchTimeZone(location)
	if err != nil {
		log.Error.Println(err)
	}

	return timezone
}
