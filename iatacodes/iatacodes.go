package iatacodes

import (
	"net/http"
	"time"
)

const (
	URL     = "http://iatacodes.org/api/v4/"
	API_KEY = "549cbc97-87f6-4763-b481-8e175677dbc9"
)

var client *http.Client

func init() {
	client = &http.Client{
		Timeout: time.Minute,
	}
}
