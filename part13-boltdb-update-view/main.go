package main

import (
	"fmt"
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

	// create table
	err = db.View(func(tx *bolt.Tx) error {
		// create BlockBucket table
		bucket := tx.Bucket([]byte("BlockBucket"))

		// put data into table
		if bucket != nil {
			data := bucket.Get([]byte("l"))
			fmt.Printf("l->\t%s\n", data)
			data = bucket.Get([]byte("ll"))
			fmt.Printf("ll->\t%s\n", data)
		}
		return nil
	})

	if err != nil {
		log.Panic(err)
	}

}
