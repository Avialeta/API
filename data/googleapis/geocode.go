package googleapis

import (
	"encoding/json"
	"errors"
	"net/url"

	"github.com/avialeta/api/log"
)

type Result struct {
	Geometry Geometry `json:"geometry"`
}

type Geometry struct {
	Location Location `json:"location"`
}

type Location struct {
	Lat float32 `json:"lat"`
	Lng float32 `json:"lng"`
}

func FetchLocation(address string) (Location, error) {
	values := url.Values{
		"address": {address},
		"key":     {API_KEY},
	}
	url := URL + "geocode/json?" + values.Encode()
	log.Info.Printf("FetchLocation: GET %s", url)

	resp, err := client.Get(url)
	if err != nil {
		return Location{}, errors.New(url + " " + err.Error())
	}
	defer resp.Body.Close()

	results := struct {
		Results []Result `json:"results"`
	}{}

	if err = json.NewDecoder(resp.Body).Decode(&results); err != nil {
		return Location{}, errors.New(url + " " + err.Error())
	}

	if len(results.Results) == 0 {
		return Location{}, errors.New(url + " len(results.Results) == 0")
	}

	location := results.Results[0].Geometry.Location

	return location, nil
}
