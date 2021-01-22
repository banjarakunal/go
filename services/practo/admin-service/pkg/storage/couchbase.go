package storage

import (
	"encoding/json"
	"log"

	"github.com/couchbase/gocb"
)

type CouchbaseDB struct {
	b *gocb.Bucket
}

type Doctor struct {
	ID   string `json:"id"`
	Name string `json:"Name"`
}

//SetupStorage setups cochbase data base and returns bucket object
func SetupStorage(bucketName string) *CouchbaseDB {
	cluster, err := gocb.Connect("http://127.0.0.1:8091")

	if err != nil {
		log.Fatalf("Error connecting:  %v", err)
	}

	bucket, err := cluster.OpenBucket(bucketName, "nbn-bucket01!")
	if err != nil {
		log.Fatalf("Error getting bucket:  %v", err)
	}

	return &CouchbaseDB{bucket}

}

func (c *CouchbaseDB) AddDoc() {

	cas, err := c.b.Upsert("dr1", Doctor{
		ID:   "1",
		Name: "Dr. Abc",
	}, 0)

	if err != nil {
		log.Fatal("Error while inserting data ", err)
	}

	log.Println("Cas for insert document", cas)

	var doctor Doctor
	c.b.Get("dr1", &doctor)

	b, _ := json.Marshal(doctor)

	log.Println(string(b))

}
