package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	mux := http.NewServeMux()
	var d DoctorDetails
	mux.Handle("/", &d)
	s := http.Server{
		Addr:         ":9090",
		Handler:      mux,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)
	log.Println("Received Terminate requet, gracefull shutdown", <-sigChan)

	tc, _ := context.WithTimeout(context.Background(), 1*time.Millisecond)
	log.Println("Server shutdown")

	s.Shutdown(tc)
}

var detailsStorage []DoctorDetails

//DoctorDetails is the doctor details to be register
type DoctorDetails struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Degree string `json:"degree"`
}

//ServeHttp server http request
func (d *DoctorDetails) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case "GET":
		w.WriteHeader(http.StatusOK)
		fmt.Println("Details :::", detailsStorage)
		err := json.NewEncoder(w).Encode(detailsStorage)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	case "POST":
		var d DoctorDetails
		err := json.NewDecoder(r.Body).Decode(&d)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		detailsStorage = append(detailsStorage, d)
		fmt.Println("Body recieved :::", d.Name)
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"message":"status accepted"}`))
	case "DELETE":
		fmt.Println("IN delete operation")
		url := r.URL
		q := url.Query()
		id := q["id"]
		fmt.Println(id[0])

		for i := range detailsStorage {
			if detailsStorage[i].ID == id[0] {
				detailsStorage = append(detailsStorage[:i], detailsStorage[i+1:]...)
			}
		}
		w.WriteHeader(http.StatusAccepted)
		err := json.NewEncoder(w).Encode(detailsStorage)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message":"Api operation not allowed"}`))

	}

}
