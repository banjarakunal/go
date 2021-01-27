package storage

import (
	"fmt"
	"log"

	"github.com/banjarakunal/go/services/go/practo/admin-services/pkg/adding"
	"github.com/google/uuid"
)

func (s *Storage) Add(d adding.Doctor) {
	log.Println("In repo")
	id := fmt.Sprintf("%s%s", "go", uuid.New().String())
	log.Println("Addding document for id", id)
	cas, err := s.b.Upsert(id, d, 0)

	if err != nil {
		log.Fatal("Error while inserting data ", err)
	}

	log.Println("Cas for insert document", cas)

	//var doctor Doctor
	//c.bucket.B.Get("dr1", &doctor)

	//b, _ := json.Marshal(doctor)

	//log.Println(string(b))

}

func (s *Storage) Get() []adding.Doctor {
	log.Println("In repo")

	var doctor []adding.Doctor
	s.b.Get("dr1", &doctor)
	return doctor

}
