package ovio

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/avialeta/api/log"
)

type Coordinates struct {
	Longitude float32 `json:"longitude"`
	Latitude  float32 `json:"latitude"`
}

type Country struct {
	Id          string      `json:"id"`
	Name        string      `json:"name"`
	Continent   string      `json:"continent"`
	Coordinates Coordinates `json:"coordinates"`
}

type City struct {
	Code    string `json:"code"`
	Country string `json:"Country"`
	Name    string `json:"name"`
}

type Airport struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	City    string `json:"city"`
	Country string `json:"country"`
}

func FetchCountries() (map[string]Country, error) {
	url := URL + "countries.json"
	log.Info.Printf("FetchCountries: GET %s", url)

	resp, err := http.Get(url)
	if err != nil {
		return nil, errors.New(url + " " + err.Error())
	}
	defer resp.Body.Close()

	countries := make(map[string]Country)
	if err = json.NewDecoder(resp.Body).Decode(&countries); err != nil {
		return nil, errors.New(url + " " + err.Error())
	}

	return countries, nil
}

func FetchCities(country string) (map[string]City, error) {
	url := URL + fmt.Sprintf("cities/%s.json", country)
	log.Info.Printf("FetchCities: GET %s", url)

	resp, err := http.Get(url)
	if err != nil {
		return nil, errors.New(url + " " + err.Error())
	}
	defer resp.Body.Close()

	cities := make(map[string]City)
	if err = json.NewDecoder(resp.Body).Decode(&cities); err != nil {
		return nil, errors.New(url + " " + err.Error())
	}

	return cities, nil
}

func FetchAirports(country string) (map[string]Airport, error) {
	url := URL + fmt.Sprintf("airports/%s.json", country)
	log.Info.Printf("FetchAirports: GET %s", url)

	resp, err := http.Get(url)
	if err != nil {
		return nil, errors.New(url + " " + err.Error())
	}
	defer resp.Body.Close()

	airports := make(map[string]Airport)
	if err = json.NewDecoder(resp.Body).Decode(&airports); err != nil {
		return nil, errors.New(url + " " + err.Error())
	}

	return airports, nil
}
