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
	err = db.Update(func(tx *bolt.Tx) error {
		// create BlockBucket table
		bucket, err := tx.CreateBucket([]byte("BlockBucket"))
		if err != nil {
			return fmt.Errorf("create bucket failed, err=%s", err)
		}
		// put data into table
		if bucket != nil {
			err := bucket.Put(
				[]byte("l"),
				[]byte("send 1 bitcoin to satoshi"),
			)
			if err != nil {
				return fmt.Errorf("put failed, err=%s", err)
			}
		}
		return nil
	})

	if err != nil {
		log.Panic(err)
	}

}
