package main

import (
	"log"

	"github.com/boltdb/bolt"
)

func main() {
	// create or open database
	db, err := bolt.Open(
		"database.db",
		0600,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

}
