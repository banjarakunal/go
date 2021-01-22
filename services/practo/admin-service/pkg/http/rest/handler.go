package rest

import (
	"net/http"
	"time"
)

//InitHandler handles the incoming request and act as controller return the http.server
func InitHandler() *http.Server {

	mux := http.NewServeMux()

	s := http.Server{
		Addr:         ":9090",
		Handler:      mux,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	return &s

}
