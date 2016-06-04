package googleapis

import (
	"net/http"
	"time"
)

const (
	API_KEY = "AIzaSyCzzhX20WSJKNqwm3D9ADq5nT0ILKkJctE"

	URL = "https://maps.googleapis.com/maps/api/"
)

var client *http.Client

func init() {
	client = &http.Client{
		Timeout: time.Minute,
	}
}
