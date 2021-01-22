package main

import (
	"github.com/banjarakunal/go/services/go/practo/admin-services/pkg/storage"
)

func main() {
	db := storage.SetupStorage("nbn-bucket01-dev")
	db.AddDoc()

}
