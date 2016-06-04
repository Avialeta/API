package api

import (
	"net/http"

	"github.com/avialeta/api/job"
	"github.com/avialeta/api/log"
)

func init() {
	mux.HandleFunc("/locations/", handleLocations)
}

func handleLocations(w http.ResponseWriter, r *http.Request) {
	locations, err := job.SearchLocations(r.URL.Query())
	if err != nil {
		log.Error.Print(err)
		handleInternalServerError(w)
		return
	}

	if locations == nil {
		handleNotFound(w)
		return
	}

	setHeader(w)
	w.Write(locations)
}
