package googleapis

import (
	"encoding/json"
	"errors"
	"net/url"
	"strconv"
	"time"

	"github.com/avialeta/api/log"
)

func FetchTimeZone(location Location) (string, error) {
	timestamp := strconv.Itoa(int(time.Now().Unix()))
	lat := strconv.FormatFloat(float64(location.Lat), 'f', 4, 32)
	lng := strconv.FormatFloat(float64(location.Lng), 'f', 4, 32)

	values := url.Values{
		"location":  {lat + "," + lng},
		"timestamp": {timestamp},
		"key":       {"AIzaSyApxatlrUI10Kp3avJsXhurUXuvf0MINuI"},
	}
	url := URL + "timezone/json?" + values.Encode()
	log.Info.Printf("FetchTimeZone: GET %s", url)

	resp, err := client.Get(url)
	if err != nil {
		return "", errors.New(url + " " + err.Error())
	}
	defer resp.Body.Close()

	res := struct {
		TimeZoneId string `json:"timeZoneId"`
	}{}

	if err = json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return "", errors.New(url + " " + err.Error())
	}

	return res.TimeZoneId, nil
}
