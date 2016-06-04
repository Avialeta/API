package ovio

import (
	"net/http"
	"os"
	"time"
)

const URL = "http://letaby.tripcloud.eu/api/"

var (
	client   *http.Client
	password = os.Getenv("TRIPCLOUD_PASSWORD")
	partner  = "letaby"
)

func init() {
	client = &http.Client{
		Timeout: time.Minute,
	}
}
