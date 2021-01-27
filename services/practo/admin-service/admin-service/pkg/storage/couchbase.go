package storage

import (
	"log"

	"github.com/couchbase/gocb"
)

type Storage struct {
	b *gocb.Bucket
}

//SetupStorage setups cochbase data base and returns bucket object
func SetupStorage(bucketName string) *Storage {
	cluster, err := gocb.Connect("http://127.0.0.1:8091")

	if err != nil {
		log.Fatalf("Error connecting:  %v", err)
	}

	bucket, err := cluster.OpenBucket(bucketName, "nbn-bucket01!")
	if err != nil {
		log.Fatalf("Error getting bucket:  %v", err)
	}

	return &Storage{bucket}

}
