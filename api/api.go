package api

import (
	"net/http"
	"strconv"
	"time"

	"github.com/avialeta/api/log"
)

const (
	ADDR = ":8080"
)

var mux = http.NewServeMux()

func Serve() {
	s := &http.Server{
		Addr:           ADDR,
		Handler:        mux,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   5 * time.Second,
		MaxHeaderBytes: http.DefaultMaxHeaderBytes,
	}

	if err := s.ListenAndServe(); err != nil {
		log.Error.Print(err)
	}
}

func setHeader(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
}

func handleNotFound(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
	w.Write(nil)
}

func handleInternalServerError(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write(nil)
}

func init() {
	mux.HandleFunc("/", handleApi)
	mux.HandleFunc("/foo/", handleFoo)
}

func handleApi(w http.ResponseWriter, r *http.Request) {
	time.Sleep(3 * time.Second)

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte("ok"))
}

func handleFoo(w http.ResponseWriter, r *http.Request) {
	v := r.URL.Query()
	if v["delay"][0] != "" {
		delay, err := strconv.Atoi(v["delay"][0])
		if err == nil {
			time.Sleep(time.Duration(delay) * time.Second)
		}
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte("ok"))
}
