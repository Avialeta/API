package iatacodes

import (
	"encoding/json"
	"net/url"

	"github.com/avialeta/api/log"
)

type Airline struct {
	Icao        string `json:"icao"`
	Code        string `json:"code"`
	Name        string `json:"name"`
	CountryCode string `json:"country_code"`
	CountryName string `json:"country_name"`
}

func FetchAirlines() ([]Airline, error) {
	params := url.Values{
		"api_key": {API_KEY},
	}

	airlines := struct {
		Response []Airline `json:"response"`
	}{}

	url := URL + "airlines?" + params.Encode()
	log.Info.Printf("FetchAirlines: GET %s", url)

	resp, err := client.Get(url)
	if err != nil {
		return airlines.Response, err
	}
	defer resp.Body.Close()

	if err = json.NewDecoder(resp.Body).Decode(&airlines); err != nil {
		return airlines.Response, err
	}

	return airlines.Response, nil
}
