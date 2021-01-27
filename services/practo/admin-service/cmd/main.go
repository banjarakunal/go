package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/banjarakunal/go/services/go/practo/admin-services/pkg/adding"
	"github.com/banjarakunal/go/services/go/practo/admin-services/pkg/http/rest"
	"github.com/banjarakunal/go/services/go/practo/admin-services/pkg/storage"

	"github.com/banjarakunal/go/services/go/practo/admin-services/pkg/retrieving"
)

func main() {

	db := storage.SetupStorage("nbn-bucket01-dev")
	//db.AddDoc()
	cs := adding.NewService(db)
	gs := retrieving.NewService(db)
	log.Println("Calling handlere")
	s := rest.InitHandler(cs, gs)

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			log.Fatal("Error while starting the server..", err)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)
	log.Println("Received Terminate requet, gracefull shutdown", <-sigChan)

	tc, _ := context.WithTimeout(context.Background(), 3000*time.Millisecond)
	log.Println("Server shutdown")

	s.Shutdown(tc)

}
