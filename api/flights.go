package api

import (
	"net/http"

	"github.com/avialeta/api/job"
)

func init() {
	mux.HandleFunc("/flights/", handleFlights)
}

func handleFlights(w http.ResponseWriter, r *http.Request) {
	flights, err := job.FetchFlights(r.URL.Query())
	if err != nil {
		handleInternalServerError(w)
		return
	}

	if flights == nil {
		handleNotFound(w)
		return
	}

	setHeader(w)
	w.Write(flights)
}
