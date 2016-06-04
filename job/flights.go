package job

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"sort"
	"strconv"
	"time"

	"github.com/avialeta/api/log"
	"github.com/avialeta/api/ovio"
	"github.com/avialeta/api/search"
)

const CONNECTION_FLIGHTS = 3

type Variants []Variant

func (v Variants) Len() int {
	return len(v)
}

func (v Variants) Less(i, j int) bool {
	return v[i].Price < v[j].Price
}
func (v Variants) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}

type Variant struct {
	Id       string    `json:"id"`
	Price    float32   `json:"price"`
	Currency string    `json:"currency"`
	Segments []Segment `json:"segments"`
}

type Segment struct {
	Flights []Flight `json:"flights"`
}

type Flight struct {
	Airline       string `json:"airline"`
	Departure     string `json:"departure"`
	DepartureDate string `json:"departureDate"`
	Arrival       string `json:"arrival"`
	ArrivalDate   string `json:"arrivalDate"`
	TimeDiff      string `json:"timeDiff"`
}

func FetchFlights(v url.Values) ([]byte, error) {
	v["pointA"][0] = search.GetCodeByName(v["pointA"][0])
	v["pointB"][0] = search.GetCodeByName(v["pointB"][0])

	if inboundDate, ok := v["inboundDate"]; !ok || len(inboundDate) == 0 {
		delete(v, "inboundDate")
	}

	rawVariants, err := ovio.FetchFlights(v)
	if err != nil {
		return nil, err
	}

	if rawVariants == nil {
		return nil, errors.New("Flights not found.")
	}

	filterRawVariants(&rawVariants)

	variants := processVariants(rawVariants)

	if len(variants) == 0 {
		return nil, nil
	}

	flights := struct {
		Variants Variants `json:"variants"`
	}{variants}

	data, err := json.Marshal(flights)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func filterRawVariants(rawVariants *ovio.Variants) {
	for id, variant := range *rawVariants {
		for i := range variant.Segments {
			if len(variant.Segments[i].Flights) > CONNECTION_FLIGHTS {
				delete(*rawVariants, id)
				break
			}
		}
	}
}

func processVariants(rawVariants ovio.Variants) []Variant {
	if len(rawVariants) == 0 {
		return nil
	}

	variants := make([]Variant, 0, len(rawVariants))

	for id, rawVariant := range rawVariants {
		variant := Variant{
			Id:       id,
			Price:    rawVariant.Price,
			Currency: rawVariant.Currency,
			Segments: processSegments(rawVariant.Segments),
		}

		variants = append(variants, variant)
	}

	return variants
}

func processSegments(rawSegments []ovio.Segment) []Segment {
	segments := make([]Segment, 0, len(rawSegments))

	for _, rawSegment := range rawSegments {
		segment := Segment{
			Flights: processFlights(rawSegment.Flights),
		}

		segments = append(segments, segment)
	}

	return segments
}

func processFlights(rawFlights []ovio.Flight) []Flight {
	flights := make([]Flight, 0, len(rawFlights))

	for i := range rawFlights {
		flight := Flight{
			Airline:       cookAirline(rawFlights[i].OperatingCarrier, rawFlights[i].MarketingCarrier),
			Departure:     cookFlightPoint(rawFlights[i].Departure),
			DepartureDate: cookDate(rawFlights[i].DepartureDate, rawFlights[i].DepartureTime, rawFlights[i].Departure),
			Arrival:       cookFlightPoint(rawFlights[i].Arrival),
			ArrivalDate:   cookDate(rawFlights[i].ArrivalDate, rawFlights[i].ArrivalTime, rawFlights[i].Arrival),
		}

		if i+1 < len(rawFlights) {
			flight.TimeDiff = processTimeDiff(rawFlights[i], rawFlights[i+1])
		}

		flights = append(flights, flight)
	}

	return flights
}

func FilterVairants(rawVariants ovio.Variants) []Variant {
	variants := Variants{}

	for id, rawVariant := range rawVariants {
		if len(rawVariant.Segments[0].Flights) <= CONNECTION_FLIGHTS {
			variant := Variant{
				Id:       id,
				Price:    rawVariant.Price,
				Currency: rawVariant.Currency,
				Segments: FilterSegments(rawVariant.Segments),
			}

			variants = append(variants, variant)
		}
	}

	sort.Sort(variants)

	return variants
}

func FilterSegments(rawSegments []ovio.Segment) []Segment {
	segments := []Segment{}

	for i := range rawSegments {
		segment := Segment{
			Flights: FilterFlights(rawSegments[i].Flights),
		}

		segments = append(segments, segment)
	}

	return segments
}

func FilterFlights(rawFlights []ovio.Flight) []Flight {
	flights := []Flight{}

	for i := range rawFlights {
		flight := Flight{
			Airline:       cookAirline(rawFlights[i].OperatingCarrier, rawFlights[i].MarketingCarrier),
			Departure:     cookFlightPoint(rawFlights[i].Departure),
			DepartureDate: cookDate(rawFlights[i].DepartureDate, rawFlights[i].DepartureTime, rawFlights[i].Departure),
			Arrival:       cookFlightPoint(rawFlights[i].Arrival),
			ArrivalDate:   cookDate(rawFlights[i].ArrivalDate, rawFlights[i].ArrivalTime, rawFlights[i].Arrival),
		}

		// Time Diff

		flights = append(flights, flight)
	}

	return flights
}

func cookAirline(code1, code2 string) string {
	if airline, ok := Airlines[code1]; ok {
		return airline.Name
	}

	if airline, ok := Airlines[code2]; ok {
		return airline.Name
	}

	return ""
}

func cookFlightPoint(code string) string {
	if airport, ok := Airports[code]; ok {
		if city, ok := Cities[airport.CityCode]; ok {
			return fmt.Sprintf("%s (%s)", city.Name, code)
		}
	}

	name := search.GetNameByCode(code)
	return fmt.Sprintf("%s (%s)", name, code)
}

// Add three hours to the date.
func cookDate(utcDate, utcTime, code string) string {
	timezone := "UTC"
	if airport, ok := Airports[code]; ok {
		timezone = airport.TimeZone
	}

	location, err := time.LoadLocation(timezone)
	if err != nil {
		fmt.Println(err)
	}

	t, err := time.Parse(time.RFC3339, utcDate+"T"+utcTime+":00+00:00")
	if err != nil {
		log.Error.Println(err)
		return utcDate + " " + utcTime
	}

	return fmt.Sprintf("%d-%02d-%02d %02d:%02d", t.In(location).Year(), t.In(location).Month(), t.In(location).Day(), t.In(location).Hour(), t.In(location).Minute())
}

func processTimeDiff(flightA, flightB ovio.Flight) string {
	timeA, err := time.Parse(time.RFC3339, flightA.ArrivalDate+"T"+flightA.ArrivalTime+":00+00:00")
	if err != nil {
		log.Error.Println(err)
		return ""
	}

	timeB, err := time.Parse(time.RFC3339, flightB.DepartureDate+"T"+flightB.DepartureTime+":00+00:00")
	if err != nil {
		log.Error.Println(err)
		return ""
	}

	diff := timeB.Sub(timeA)
	hours := int(diff.Hours())
	minutes := int(diff.Minutes()) - hours*60

	return strconv.Itoa(hours) + "h " + strconv.Itoa(minutes) + "m"
}
