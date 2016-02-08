package ovio

import (
	"encoding/json"
	"net/url"

	"github.com/avialeta/api/log"
)

type Variants map[string]Variant

type Variant struct {
	Price    float32   `json:"price"`
	Currency string    `json:"currency"`
	Segments []Segment `json:"segments"`
}

type Segment struct {
	Flights []Flight `json:"flights"`
}

type Flight struct {
	OperatingCarrier string `json:"operatingCarrier"`
	MarketingCarrier string `json:"marketingCarrier"`
	Departure        string `json:"departure"`
	DepartureDate    string `json:"departureDate"`
	DepartureTime    string `json:"departureTime"`
	Arrival          string `json:"arrival"`
	ArrivalDate      string `json:"arrivalDate"`
	ArrivalTime      string `json:"arrivalTime"`
}

func FetchFlights(values url.Values) (Variants, error) {
	params := url.Values{
		"password": {password},
		// TODO: Remove
		"adultCount": {"1"},
	}
	for name, value := range values {
		params[name] = value
	}

	url := URL + "flights/search.json?" + params.Encode()
	log.Info.Printf("FetchCountries: GET %s", url)

	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	flights := struct {
		Variants Variants `json:"variants"`
	}{}

	if err = json.NewDecoder(resp.Body).Decode(&flights); err != nil {
		return nil, err
	}

	return flights.Variants, nil
}
