package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/banjarakunal/go/services/go/practo/admin-services/pkg/adding"
	"github.com/banjarakunal/go/services/go/practo/admin-services/pkg/retrieving"
)

//InitHandler handles the incoming request and act as controller return the http.server
func InitHandler(cs adding.Service, gs retrieving.Service) *http.Server {

	mux := http.NewServeMux()
	mux.HandleFunc("/api/doctor", adminHandler(cs, gs))

	s := http.Server{
		Addr:         ":9090",
		Handler:      mux,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	return &s

}

func adminHandler(cs adding.Service, gs retrieving.Service) func(rw http.ResponseWriter, r *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", "application/json")

		switch r.Method {
		case "GET":
			rw.WriteHeader(http.StatusOK)
			d := gs.GetDocotsDetails()
			fmt.Println("Details :::", d)
			err := json.NewEncoder(rw).Encode(d)
			if err != nil {
				http.Error(rw, err.Error(), http.StatusBadRequest)
				return
			}

		case "POST":
			var d adding.Doctor
			err := json.NewDecoder(r.Body).Decode(&d)

			if err != nil {
				http.Error(rw, err.Error(), http.StatusBadRequest)
				return
			}
			cs.AddDoctor(d)
			fmt.Println("In handler adding document")
			rw.WriteHeader(http.StatusCreated)
			rw.Write([]byte(`{"message":"status accepted"}`))

		default:
			rw.WriteHeader(http.StatusNotFound)
			rw.Write([]byte(`{"message":"Api operation not allowed"}`))
		}
	}
}
